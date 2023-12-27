package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleSSH(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("centos"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "10.0.0.20:22", config)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()
	stdinPipe, err := session.StdinPipe()
	if err != nil {
		log.Println("stdin err: ", err)
		return
	}
	stdoutPipe, err := session.StdoutPipe()
	if err != nil {
		log.Println("stdout err: ", err)
		return
	}
	stderrPipe, err := session.StderrPipe()
	if err != nil {
		log.Println("stderr err: ", err)
		return
	}
	err = session.Shell()
	if err != nil {
		log.Println("shell err:", err)
	}

	in, out := make(chan []byte), make(chan []byte)
	go func() {
		for {
			messageType, bytes, err2 := conn.ReadMessage()
			if err2 != nil {
				log.Println(err2)
				break
			}
			log.Println(messageType, err2)
			in <- bytes
		}
	}()
	go func() {
		for {
			select {
			case msg := <-in:
				log.Println("获取msg：", string(msg))
				n, err2 := stdinPipe.Write(msg)
				log.Println("写入 ", n, err2)
			}
		}
	}()
	outMsg := make([]byte, 4096)
	go func() {
		for {
			n, err2 := stdoutPipe.Read(outMsg)
			log.Println(n)
			if err2 != nil {
				log.Println("read err: ", err2)
				continue
			}
			out <- outMsg[:n]
		}
	}()
	go func() {
		for {
			n, err2 := stderrPipe.Read(outMsg)
			log.Println(n)
			if err2 != nil {
				log.Println("err err: ", err2)
				continue
			}
			out <- outMsg[:n]
		}
	}()
	go func() {
		for {
			msg := <-out
			err := conn.WriteMessage(1, msg)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	err = session.Wait()
	if err != nil {
		log.Println("wait err: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/ssh", handleSSH)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

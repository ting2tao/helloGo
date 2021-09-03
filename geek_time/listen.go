package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	server := http.Server{Addr: ":8080"}

	go func() {
		server.ListenAndServe()
	}()
	fmt.Println("hh")
	<-quit

	server.Shutdown(context.TODO())
	fmt.Println("77")
}

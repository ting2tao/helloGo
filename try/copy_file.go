package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {

	exPath, _ := os.Getwd()

	fmt.Println("exPath", exPath)
	w, _ := CopyFile(exPath+"/try/c3.txt", exPath+"/try/copy.txt")
	fmt.Sprintln(w)
}

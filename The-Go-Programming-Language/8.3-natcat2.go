package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr := os.Args[1]
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy2(os.Stdout, conn)
	go mustCopy2(os.Stdout, conn)
}

func mustCopy2(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

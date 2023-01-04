package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//var data [1024]byte
		//n, err := conn.Read(data[0:])
		//s := string(data[0:n])
		conn.Close()
		if err != nil {
			fmt.Println(err)
		}

		//fmt.Fprintf(conn, "receive: ")
		//handleConn(conn)
		// 注意 加不加 go 关键字的区别
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

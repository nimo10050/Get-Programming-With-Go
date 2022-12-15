package main

import (
	"flag"
	"fmt"
)

func main() {
	var port = flag.Int("port", 1, "test")
	var ip = flag.String("ip", "127.0.0.1", "test")
	flag.Parse()
	fmt.Println(&port)
	fmt.Println(*port)
	fmt.Println(*ip)
	//fmt.Println(&port)
}

package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[0:] {
		fmt.Println("index: ", index, "argv: ", arg)
		//resp, err := http.Get(arg)
		//if err != nil {
		//	fmt.Println("network err: ", err)
		//}
		//b, err := io.ReadAll(resp.Body)
		//resp.Body.Close()
		//fmt.Printf("%s", b)
	}

}

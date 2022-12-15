package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		// fmt.Println(i)
		go func() {
			fmt.Println("current index: ", i)
		}()
		//fmt.Println("======")
	}

	//for {
	// fmt.Printf("main goroutine: i ")
	time.Sleep(3 * time.Second)
	//}
}

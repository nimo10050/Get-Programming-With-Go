package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	for i := 0; i < 10; i++ {
		once.Do(func() {
			fmt.Println("hello world")
		})
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go myFunction(int64(i))
	}

	time.Sleep(3 * time.Second)

}

func myFunction(i int64) {
	fmt.Println(i)
}

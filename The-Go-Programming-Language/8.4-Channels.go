package main

import "fmt"

// A channel is a communication mechanism that lets one goroutine s end values to another goroutine
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	defer close(naturals)
	go func() {
		for i := 0; i < 10; i++ {
			naturals <- i
		}
		//close(naturals)
	}()

	//defer close(squares)

	go func() {
		for i := 0; i < 10; i++ {
			j := <-naturals
			squares <- j * j
		}
		//close(squares)
	}()

	//for x := range squares {
	//	fmt.Println(x)
	//}

	for {
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
	}

}

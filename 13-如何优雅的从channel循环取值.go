package main

import "fmt"

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	// 向 ch1 写值
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}

		close(ch1)

	}()

	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * 2
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}

}

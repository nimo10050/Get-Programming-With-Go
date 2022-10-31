package main

import "fmt"

// 让 10 个 routine 按顺序执行
func main() {

	fmt.Println("begin.")

	ch := make(chan int, 4)

	for i := 0; i < 4; i++ {
		go func() {
			fmt.Println("seq: ", i)
			ch <- i
		}()
	}

	for j := 0; j < 4; j++ {
		var t = <-ch
		fmt.Println(t)
	}

	fmt.Println("end.")

}

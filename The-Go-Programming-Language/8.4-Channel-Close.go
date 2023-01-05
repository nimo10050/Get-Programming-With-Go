package main

import "fmt"

// 通道关闭的小例子
func main() {

	// 启动一个 goroutine 往通道中发送数据
	ch := make(chan int)
	go Square(ch)

	// 阻塞, 直到通道关闭
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("channel is closed.")
			return
		} else {
			fmt.Println("receive : ", v)
		}
	}
}

// Square 计算平方值
func Square(ch chan int) {
	for i := 0; i < 9; i++ {
		ch <- i * i
	}

	close(ch)
}

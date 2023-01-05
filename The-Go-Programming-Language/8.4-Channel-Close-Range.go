package main

import "fmt"

// 通道关闭的小例子
func main() {

	// 启动一个 goroutine 往通道中发送数据
	ch := make(chan int)
	go Square(ch)

	// 如果通道关闭， range 会知道的
	for v := range ch {
		fmt.Println("receive : ", v)
	}
}

// Square 计算平方值
func Square(ch chan int) {
	for i := 0; i < 9; i++ {
		ch <- i * i
	}

	// 如果这里不 close， fatal error: all goroutines are asleep - deadlock!
	close(ch)
}

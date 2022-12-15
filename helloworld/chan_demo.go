package main

import "fmt"

//chan 管道
func main() {
	// 声明一个通道, 可以使用 make 函数声明初始化，比如下面， 我们只能往通道里面丢 int 类型的数据
	ch := make(chan int)
	go func() {
		var sum int = 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum // 发送到通道
	}()


	// 在计算 sum 和 goroutine 没有执行完， 把值赋给 ch 通过之前， 下面的语句会一直等待。
	x := <- ch // 从通道里面读取值， 并把读取的值赋值给 x
	fmt.Println(x)
	close(ch)// 内置 close 函数关闭


	// 单向通道
	//var send chan <- int // 只能发送
	//var receive <-chan int // 只能接收
}

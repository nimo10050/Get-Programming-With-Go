package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 向 chan1 写入数据
	ch1 := make(chan int, 1)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 1
	}()

	// 2. 向 chan2 写入数据
	ch2 := make(chan int, 2)
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	// select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句
	select {
	case <-ch1:
		fmt.Println("read data from ch1")
	case <-ch2:
		fmt.Println("read data from ch2")
	}

}

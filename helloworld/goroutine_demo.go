package main

import (
	"fmt"
	"sync"
)

func main() {
	// 有点像 java 中的 CountDownLatch
	var wg sync.WaitGroup
	wg.Add(2)

	// 创建一个 goroutine 是通过 go 关键字
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			fmt.Println("A thread: ", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			fmt.Println("B thread: ", i)
		}
	}()

	wg.Wait()
}
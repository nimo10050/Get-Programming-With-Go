package main

import (
	"fmt"
	"sync"
)

// 定义一个共享资源
var count = 0

// 修改 10000 次
var times = 10000

var sw sync.WaitGroup

func main() {

	fmt.Println("修改之前的 count = ", count)
	sw.Add(times)
	// 定义 10 个 goRoutine 并发更新 count 变量
	for i := 0; i < times; i++ {
		go incrementCount()
	}

	sw.Wait() // time.Sleep(10 * time.Second)

	fmt.Println("修改完后的 count = ", count)
}

// + 1 操作
func incrementCount() {
	count++
	sw.Done()
}

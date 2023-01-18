package main

import (
	"context"
	"fmt"
	"time"
)

// context 的简单使用
func main() {
	// 创建一个 context， 1 s 后， 自动结束
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()

	// 开启一个 routine 去执行任务
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Other go routine Done")
		case <-time.After(time.Millisecond * 500):
			fmt.Println("Other go routine consume 500 ms")
		}
	}()

	// 等待 1 s
	select {
	case <-ctx.Done():
		fmt.Println("Main go routine Done")
	}
}

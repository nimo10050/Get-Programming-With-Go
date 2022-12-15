package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				fmt.Println("program exit..")
				return
			default:
				fmt.Println("monitor.....")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("notify program stop.")
	cancel()
	// 为了检测监控是否停止， 如果没有监控输出， 就表示停止了
	time.Sleep( 1 * time.Second)
}

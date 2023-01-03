package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibB := fib(n)
	fmt.Println("\rFibonacci(%d) = %d\n", n, fibB)
}

// 这尼玛竟然只是一个动画 haha
func spinner(delay time.Duration) {
	for {
		// 控制台循环播放 -\|/
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

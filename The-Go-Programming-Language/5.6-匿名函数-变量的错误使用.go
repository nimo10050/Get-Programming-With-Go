package main

import (
	"fmt"
)

func main() {
	var rmdirs []func()
	// 变量的作用域
	for _, dir := range tempDirs() {
		fmt.Println("Make dir: ", dir)
		// 共享 dir 变量
		rmdirs = append(rmdirs, func() {
			fmt.Println("Romove dir: ", dir)
		})
	}

	// 这里 执行 rmdir 函数时， 函数里拿到的 dir 其实是 tempDirs[末位]
	for _, rmdir := range rmdirs {
		rmdir()
	}
}

func tempDirs() []string {
	return []string{"123", "456", "789"}
}

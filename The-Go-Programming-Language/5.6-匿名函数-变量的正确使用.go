package main

import (
	"fmt"
)

func main() {
	var rmdirs []func()
	// 变量的作用域
	for _, dir := range tempDirs() {
		// 这里用个临时变量， 直接就解决了变量共享问题
		dir := dir
		fmt.Println("Make dir: ", dir)
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

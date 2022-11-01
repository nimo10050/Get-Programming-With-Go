package main

import "fmt"

func main() {
	var name string = "abc"

	// 取内存地址
	temp := &name

	fmt.Println(temp)

	// 取地址中的值
	fmt.Println(*temp)
}

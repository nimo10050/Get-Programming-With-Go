package main

import "fmt"

var littleMap = make(map[string]string)

func main() {
	littleMap["zhangsan"] = "man"
	// 初始化变量
	if name, ok := littleMap["zhangsan"]; ok {
		fmt.Println(name)
	}
}

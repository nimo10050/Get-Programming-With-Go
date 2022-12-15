package main

import "fmt"

// * 和 & 的作用
// 取值 : & 取变量内存地址， * 取内存地址中的值
// 存值 : * 存入地址
func main() {
	var name string = "abc"

	// 取内存地址
	temp := &name

	fmt.Println(temp)

	// 取地址中的值
	fmt.Println(*temp)


	// 声明一下 m 的value 只能存对象
	m := make(map[string]Value)
	// 声明一下 m1 的value 只能存地址
	m1 := make(map[string]*Value)

	val := Value{Name: "zhangsan"}
	m["key"] = val
	// 存入地址
	m1["key"] = &val

	fmt.Println(m["key"])
	fmt.Println(m1["key"])

	fmt.Println("修改后")

	val = Value{Name: "wangwu"}
	fmt.Println(m["key"])
	fmt.Println(m1["key"])
}

type Value struct {
	Name string
}


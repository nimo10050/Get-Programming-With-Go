package main

import "fmt"

// go 语言中如何声明变量

// (前提): go 语言中的变量分为全局变量和局部变量
//
// 1. 使用 var 关键字 比如: var s = "abc",或者 var s string = "abc"

// 2. 简写 s := "abc",

// 3. 使用 var 关键字时, 如果不赋值, 需要声明类型，例如: var s string

// (注意): 声明全局变量 需要用 var 关键字， := 只能用在函数体内部。

var i = 1

func main() {
	var s = "aaa"
	var ss string = "bbb"
	sss := "ccc"

	fmt.Println("s=", s, ", ss=", ss, ", sss=", sss)

	var ssss string

	// var sssss 这样写是错误的， 必须赋值

	// 如果变量没使用， 可以使用这种方式丢弃
	_ = ssss

}

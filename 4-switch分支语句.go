package main

import "fmt"

func main() {
	// Go里面switch默认相当于每个case最后带有break，
	// 匹配成功后不会自动向下执行其他case，而是跳出整个switch,
	// 但是可以使用fallthrough强制执行后面的case代码
	var command = "get it"
	switch command {
	case "get it":
		fmt.Println("bingo")
		fallthrough
	case "not done":
		fmt.Println("try it again")
	default:
		fmt.Println("bye bye")
	}

}

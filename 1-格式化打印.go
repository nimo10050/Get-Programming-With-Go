package main

import "fmt"

func main() {
	// -15 表示与后面间隔 15 个 单位
	//  15 表示与前面间隔 15 个 单位
	fmt.Printf("%-15v $%4v \n", "SpaceX", "15")
	fmt.Printf("%15v $%4v \n", "SpaceX", "15")

	// 小数格式化 109.2f   109 是 109 个间隔位， 2 两位小数
	fmt.Printf("%109.2f", 1234.789)
}

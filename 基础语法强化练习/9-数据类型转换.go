package main

import "fmt"

func main() {
	var i int64 = 100
	var f = float64(i)
	fmt.Println("convert int 2 float: ", f)

	// 报错
	// var s string = "1"
	// i = int64(s)
	// f = float64(s)

	fmt.Println("conver string 2 float: ", f)

}

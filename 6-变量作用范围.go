package main

import (
	"fmt"
)

func main() {
	var count = 0
	//for count < 10 {
	//	//var num = rand.Intn(10) + 1
	//	//fmt.Println(num)
	//	count++
	//}
	//
	//// result: 10
	//fmt.Println(count)

	for count = 100; count > 10; count-- {
		fmt.Println(count)
	}
	fmt.Println("result:", count)
}

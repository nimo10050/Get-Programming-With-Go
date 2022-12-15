package main

import "fmt"

func main() {
	// for 死循环
	var count = 0
	for {
		count++
		if count > 30 {
			fmt.Println("count: ", count)
			break
		}
	}

	for count > 0 {
		fmt.Println(count)
		count--
	}
}

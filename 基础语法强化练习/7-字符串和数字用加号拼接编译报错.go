package main

import "fmt"

func main() {
	var num int = 10
	var s string = "abc"
	// Invalid operation: s + num (mismatched types string and int)
	// var ss = s + num

	// avoid error message
	fmt.Println(num, s)
}

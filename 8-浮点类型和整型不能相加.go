package main

import "fmt"

func main() {
	var f float64 = 6.444
	var i int64 = 100

	// Invalid operation: f + i (mismatched types float64 and int64)
	// var ff float64 = f + i

	fmt.Println(f, i)
}

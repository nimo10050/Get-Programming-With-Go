package main

import "fmt"

func main() {
	var command = "get it"

	if command == "get it" {
		fmt.Println("bingo")
	} else if command == "not done" {
		fmt.Println("try it again")
	} else {
		fmt.Println("bye bye")
	}
}

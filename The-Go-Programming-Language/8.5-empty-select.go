package main

import "fmt"

func service() {
	fmt.Println("Hello from service ")
}

func main() {
	fmt.Println("main() started")

	go service()

	// 這個 select 會永遠 block 在這
	// 最终： fatal error: all goroutines are asleep - deadlock!
	select {}

	fmt.Println("main() stopped")
}

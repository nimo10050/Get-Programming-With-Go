package main

import (
	"fmt"
	"sync"
)

func main() {
	var num = 6
	var wg sync.WaitGroup

	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		wg.Add(1)
		go Worker(&wg, ch, i)
	}

	wg.Wait()

	var numbers []int
	for i := 0; i < num; i++ {
		numbers = append(numbers, <-ch)
	}
	fmt.Println(numbers)

	//defer close(ch)

}

func Worker(wg *sync.WaitGroup, ch chan int, i int) {

	// 写入通道
	ch <- i
	defer wg.Done()
}

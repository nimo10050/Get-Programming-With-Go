package main

import (
	"fmt"
	"sync"
	"time"
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
	// 如果是在 main goroutine 中使用 for val := range channel {} 的寫法時，最後 channel 沒有被 close 的話程式會 deadlock。
	// 但如果是在其他的 goroutine 中使用，即使沒有 close 也不會 deadlock，但為了不必要的 bug 產生，一般都還是將其關閉。
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}

	// 在其他 routine 中 range
	//go func() {
	//	for v := range ch {
	//		fmt.Println(v)
	//	}
	//}()

	time.Sleep(3 * time.Second)

}

func Worker(wg *sync.WaitGroup, ch chan int, i int) {

	// 写入通道
	ch <- i
	defer wg.Done()
}

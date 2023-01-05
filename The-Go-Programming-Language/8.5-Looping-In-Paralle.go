package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func makeThumbnails(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()

		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func ImageFile(filename string) (string, error) {
	if filename == "ddd" {
		fmt.Println("DDDD")
		return "", fmt.Errorf("error 拉")
	}
	return "", nil
}

func main() {
	ss := []string{"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def"}
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
	//"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def"}
	err := makeThumbnails4(ss)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("over")
}

func makeThumbnails4(filenames []string) error {
	// 定义一个无缓冲通道
	errors := make(chan error)
	for _, f := range filenames {
		// worker
		go func(f string) {
			fmt.Println("scale  ", f)
			_, err := ImageFile(f)
			fmt.Println("开始阻塞了...")
			// 由于接收方无法接收消息, 所以这里只能等待
			errors <- err
			fmt.Println("解除阻塞了...")
		}(f)
	}

	for range filenames {
		// 这里有个小问题, 因为我们一开始定义的是一个无缓冲通道
		// 当此处从通道中接收到第一个 非 nil 的 error 时，这里直接 return 了, 导致接收方没人接收数据了
		// 但是上面的 goroutine 还在排队往通道发送数据，最终导致上面的 goroutine 就阻塞住了
		if err := <-errors; err != nil {
			fmt.Println("yes error")
			return err
		} else {
			fmt.Println("NO ERROR")
		}
	}

	return nil
}

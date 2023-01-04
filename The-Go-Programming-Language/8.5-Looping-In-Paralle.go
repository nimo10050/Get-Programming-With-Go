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
	ss := []string{"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def",
		"ddd", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def", "def", "ddd", "def", "ddd", "def"}
	err := makeThumbnails4(ss)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("over")
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error, 50)
	for _, f := range filenames {
		// worker

		go func(f string) {
			fmt.Println("scale  ", f)
			_, err := ImageFile(f)
			errors <- err
			fmt.Println("SEND OVER")
		}(f)
	}

	for range filenames {
		fmt.Println("=====")
		if err := <-errors; err != nil {
			fmt.Println("yes error")
			return err
		} else {
			fmt.Println("NO ERROR")
		}
	}

	return nil
}

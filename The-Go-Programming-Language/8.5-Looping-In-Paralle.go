package main

import (
	"fmt"
	"log"
	"os"
	"sync"
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
	ss := []string{"ddd", "def", "ddd", "def", "ddd", "def"}
	err := makeThumbnails4(ss)
	if err != nil {
		fmt.Println(err)
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		// worker

		go func(f string) {
			fmt.Println("scale  ", f)
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			fmt.Println("yes error")
			return err
		} else {
			fmt.Println("no error")
		}
	}
	return nil
}

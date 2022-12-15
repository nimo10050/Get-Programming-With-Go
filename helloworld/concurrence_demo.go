package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count int32
	wg sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	// chan
	wg.Add(2)
	// go incCount()
	// go incCount()
	go incCountSafeByMutex()
	go incCountSafeByMutex()
	wg.Wait()
	fmt.Println(count)
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		// 暂停
		runtime.Gosched()
		value ++
		count = value
	}
}

// 线程安全
func incCountSafeByAtomic() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := atomic.LoadInt32(&count)
		// 暂停
		runtime.Gosched()
		value ++
		atomic.StoreInt32(&count, value)
	}
}

func incCountSafeByMutex() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		value := count
		// 暂停
		runtime.Gosched()
		value ++
		count = value
		mutex.Unlock()
	}
}
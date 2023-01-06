package main

import "fmt"

/* Go 語言使用 Select 四大用法：https://blog.wu-boy.com/2019/11/four-tips-with-select-in-golang/ */

func main() {
	// STEP 1：建立一個只能裝 buffer size 為 1 資料
	ch := make(chan int, 1)
	ch <- 1

	select {
	case ch <- 2:
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)
	default:
		// ch 中的內容超過 1 時，但若把 channel buffer size 的容量改成 2，就不會走到 default
		fmt.Println("channel blocking")
	}
}

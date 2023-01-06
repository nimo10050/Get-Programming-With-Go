package main

import "fmt"

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go service1(ch1)
	go service2(ch2)

	// 如果所有的 case 都沒有接收到 channel 傳來的資料，那麼 select 會一直阻塞（block）在那，
	// 直到有任何的 case 收到資料後（unblock）才會繼續執行

	// 如果同一時間有多個 case 收到 channel 傳來的資料（有多個 channel 同時 non-blocking），
	// 那個會從所有這些 non-blocking 的 cases 中隨機挑選一個，接著才繼續執行
	select {
	case <-ch1:
		fmt.Println("hello ch1")
	case <-ch2:
		fmt.Println("hello ch2")
	}
}

func service2(ch2 chan string) {
	fmt.Println("send hello to ch2")
	ch2 <- "hello chh2"
}

func service1(ch1 chan string) {
	fmt.Println("send hello to ch1")
	ch1 <- "hello chh1"
}

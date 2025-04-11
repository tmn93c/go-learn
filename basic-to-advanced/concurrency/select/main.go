package main

import (
	"fmt"
	"time"
)

// select được sử dụng để xử lý nhiều channel đồng thời.
// Giải thích:

// select lắng nghe nhiều channel cùng lúc.
// Khi một channel sẵn sàng (có dữ liệu), case tương ứng sẽ thực thi.
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "Message from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println(msg1)
		case msg2 := <-channel2:
			fmt.Println(msg2)
		}
	}
}

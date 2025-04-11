package main

import "fmt"

// Channels mặc định là unbuffered (không có bộ đệm), nghĩa là việc gửi và nhận phải xảy ra đồng thời. Buffered channels cho phép lưu trữ một số lượng phần tử nhất định.
func main() {
	channel := make(chan int, 2) // Tạo buffered channel với kích thước 2

	channel <- 1
	channel <- 2

	fmt.Println(<-channel) // Nhận giá trị 1
	fmt.Println(<-channel) // Nhận giá trị 2
}

// Lưu ý: Nếu bạn gửi dữ liệu vào một buffered channel đã đầy hoặc nhận từ một channel trống, chương trình sẽ block (dừng lại) cho đến khi có dữ liệu phù hợp.

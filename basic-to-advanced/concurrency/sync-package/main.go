package main

import (
	"fmt"
	"sync"
)

// sync.WaitGroup: đợi nhiều goroutine hoàn tất
// sync.Mutex: tránh race condition

// func main() {
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	count := 0

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			mu.Lock()
// 			count++
// 			mu.Unlock()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(count)
// }

var once sync.Once

func initOnce() {
	fmt.Println("Initialization done!")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			once.Do(initOnce) // Chỉ thực hiện initOnce một lần
			fmt.Println("Goroutine", n, "is running")
		}(i)
	}

	wg.Wait() // Đợi cho tất cả goroutines hoàn thành
}

// Chỉ Thực Hiện Một Lần: Hàm initOnce chỉ được thực hiện một lần, bất kể số lần once.Do(initOnce) được gọi từ các goroutines khác nhau.
// Đảm Bảo Tính Đồng Bộ: Điều này rất hữu ích trong việc khởi tạo tài nguyên chỉ một lần trong một ứng dụng đa luồng.

// =>
// Sử dụng sync.Mutex để bảo vệ các vùng dữ liệu chia sẻ và đảm bảo tính chính xác trong các ứng dụng đồng thời.
// Sử dụng sync.Once để thực hiện các khởi tạo chỉ một lần, ngay cả khi có nhiều goroutines gọi đến.

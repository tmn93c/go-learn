package main

import (
	"context"
	"fmt"
	"time"
)

// Gói context được sử dụng để quản lý thời gian sống của goroutines, đặc biệt khi cần dừng chúng một cách an toàn.
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("Main done")
}

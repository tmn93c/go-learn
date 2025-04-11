package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulate fetching data from a URL
func fetchData(url string) string {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate network delay
	return fmt.Sprintf("Data from %s", url)
}

// Worker function (Fan-Out)
func worker(id int, urls <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		fmt.Printf("Worker %d fetching %s\n", id, url)
		data := fetchData(url)
		results <- data
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	urls := []string{
		"https://site1.com",
		"https://site2.com",
		"https://site3.com",
		"https://site4.com",
		"https://site5.com",
		"https://site6.com",
	}

	urlChan := make(chan string, len(urls))    // Channel for URLs (Fan-Out)
	resultChan := make(chan string, len(urls)) // Channel for results (Fan-In)

	// 	Fan-out: nhiều goroutine xử lý dữ liệu từ cùng một channel.

	// Fan-in: nhiều goroutine gửi vào cùng một channel để gom dữ liệu lại.

	// Fan-Out: Start worker goroutines
	numWorkers := 3
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, urlChan, resultChan, &wg)
	}

	// Send URLs to urlChan
	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan) // Close the channel to signal no more jobs

	// Wait for all workers to finish
	wg.Wait()
	close(resultChan) // Close the result channel

	// Fan-In: Collect results from resultChan
	for result := range resultChan {
		fmt.Println(result)
	}
}

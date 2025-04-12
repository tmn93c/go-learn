package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func slowFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return slowFibonacci(n-1) + slowFibonacci(n-2)
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		http.Error(w, "Invalid n", http.StatusBadRequest)
		return
	}

	result := slowFibonacci(n)
	fmt.Fprintf(w, "Fibonacci(%d) = %d\n", n, result)
}

func main() {
	// Pprof server chạy tại :6060
	go func() {
		log.Println("pprof active at :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

	http.HandleFunc("/fib", fibHandler)

	log.Println("App running at :8080")
	http.ListenAndServe(":8080", nil)
}

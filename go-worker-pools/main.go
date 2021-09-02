package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)

	for i := 0; i < 11; i++ {
		jobs <- i
	}
	close(jobs)

	// for j := 0; j < 100; j++ {
	// 	fmt.Println(<-results)
	// }

	for msg := range results {
		fmt.Println(msg)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	// specifying that we'll recive on the jobs channel and send on the results channe to reduce the chance of errors
	for n := range jobs {
		results <- fib(n)
	}
	close(results)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

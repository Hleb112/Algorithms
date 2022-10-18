package main

import (
	"fmt"
)

func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 50)
	results := make(chan int, 50)

	for i := 0; i < 3; i++ {
		go worker(jobs, results)
	}

	for j := 0; j < 100; j++ {
		jobs <- j
	}

	close(jobs)

	for i := 0; i < 100; i++ {
		result := <-results
		fmt.Println(result)
	}
}

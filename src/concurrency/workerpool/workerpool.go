package main

import "fmt"

const numberOfWorkers int = 8

func fib(num int) int {
	if num <= 1 {
		return num
	}

	return fib(num-1) + fib(num-2)
}

func worker(jobs <-chan int, result chan<- int) {
	for n := range jobs {
		result <- fib(n)
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for i := 0; i < numberOfWorkers; i++ {
		go worker(jobs, result)
	}

	for i := 1; i <= 100; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-result)
	}
}

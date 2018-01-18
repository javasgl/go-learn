package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for i := 0; i <= 9; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j <= 9; j++ {
		<-results
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

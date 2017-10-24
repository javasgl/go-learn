package main

import (
	"fmt"
	"time"
)

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {

		for {

			jobId, hasMore := <-jobs
			if hasMore {
				fmt.Println("recevied job:", jobId)
				time.Sleep(2 * time.Second)
			} else {
				fmt.Println("all jobs recevied")
				done <- true
				return
			}
		}
	}()

	for i := 0; i <= 10; i++ {
		jobs <- i
		fmt.Println("send job:", i)
	}
	close(jobs)
	fmt.Println("send all job")
	<-done
}

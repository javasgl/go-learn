package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 500)

	for r := range requests {
		<-limiter
		fmt.Println("request:", r, time.Now())
	}

	fmt.Println("bursty request")

	burstyLimiter := make(chan time.Time, 3)

	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for r := range burstyRequests {
		<-burstyLimiter
		fmt.Println("requset:", r, time.Now())
	}

}

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	consumer(ch)
	producer(ch)
}

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func consumer(ch <-chan int) {

	select {
	case v := <-ch:
		fmt.Println(v)
	default:
		time.Sleep(1000 * time.Millisecond)
	}
}

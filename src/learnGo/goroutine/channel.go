package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	go do1(ch)

	fmt.Println(runtime.NumCPU())
	select {
	case <-ch:
		fmt.Println("done!")
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
		close(ch)
	}
	fmt.Println(runtime.NumCPU())
}

func do1(ch chan int) {
	go do2(ch)
}

func do2(ch chan int) {
	time.Sleep(3 * time.Second)
	go func(ch chan int) {
		ch <- 1
	}(ch)
}

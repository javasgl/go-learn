package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		fmt.Println("1")
		ch <- 1
		fmt.Println("2")
	}()

	go func() {
		fmt.Println("3")
		<-ch
		fmt.Println("4")
	}()

	time.Sleep(1 * time.Second)

}

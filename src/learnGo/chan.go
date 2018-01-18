package main

import (
	"fmt"
)

var c1 = make(chan int)
var c2 = make(chan int)

func main() {

	go func() {
		fmt.Println("1111")
		c1 <- <-c2
	}()

	go func() {
		c1 <- <-c2
		fmt.Println("2222")
	}()
	<-c1
}

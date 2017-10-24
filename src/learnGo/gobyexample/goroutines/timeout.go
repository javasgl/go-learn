package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "done"
	}()

	select {
	case m := <-c1:
		fmt.Println(m)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1s")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "done2"
	}()

	select {
	case m := <-c2:
		fmt.Println(m)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}

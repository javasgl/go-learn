package main

import (
	"fmt"
	"time"
)

// 对nil的channel进行读写,均会阻塞
func main() {

	c := make(chan int, 2)
	t := time.After(time.Second)

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println("read from c", v)
			case <-t:
				c = nil
				fmt.Println("set c to  nil")
				c <- 1
				fmt.Println("sen v to nil chan")
			}
		}
	}()
	c <- 2

	time.Sleep(5 * time.Second)
}

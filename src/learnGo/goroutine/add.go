package main

import (
	"fmt"
	"time"
)

func Add(x, y int) {
	fmt.Println(x + y)
}
func main() {
	for i := 0; i <= 10; i++ {
		go Add(i, i)
	}
	time.Sleep(time.Millisecond) //wait goroutine to finish
}

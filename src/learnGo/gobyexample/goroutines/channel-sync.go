package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool, 1)
	// done := make(chan bool)

	go worker(done)

	<-done
}

func worker(done chan bool) {

	fmt.Println("start to working...")

	time.Sleep(2 * time.Second)

	fmt.Println("done")

	done <- true

}

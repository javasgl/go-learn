package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}

	limiter := make(chan bool, 10)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		limiter <- true
		go download(i, limiter, wg)
	}
	wg.Wait()
}

func download(index int, limiter chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("start to download :", index)
	time.Sleep(1 * time.Second)
	<-limiter
}

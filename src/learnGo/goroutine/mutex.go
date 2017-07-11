package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg         sync.WaitGroup
	counter    int
	counter_v1 int
	mutex      sync.Mutex
)

func main() {

	wg.Add(4)

	go incr_v1()
	go incr_v1()

	go incr()
	go incr()

	wg.Wait()
	fmt.Println("counter:", counter)
	fmt.Println("counter_v1:", counter_v1)
}

func incr_v1() {
	defer wg.Done()
	tmp := counter_v1
	runtime.Gosched()
	tmp++
	counter_v1 = tmp
}

func incr() {

	defer wg.Done()

	mutex.Lock()
	{
		tmp := counter
		runtime.Gosched()
		tmp++
		counter = tmp
	}
	mutex.Unlock()
}

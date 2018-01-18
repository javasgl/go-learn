package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops uint64 = 0

	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()

				atomic.AddUint64(&ops, 1)
				runtime.Gosched()

			}

		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)
			mutex.Lock()
			state[key] = val
			mutex.Unlock()

			atomic.AddUint64(&ops, 1)
			runtime.Gosched()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("ops:", atomic.LoadUint64(&ops))

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

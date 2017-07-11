package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {

	fmt.Println(runtime.NumCPU())

	runtime.GOMAXPROCS(4)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c", char)
		}
	}()
	go func() {
		defer wg.Done()
		for char := 'A'; char < 'A'+26; char++ {
			fmt.Printf("%c", char)
		}
	}()

	wg.Wait()

	fmt.Println("\nfinished main")
}

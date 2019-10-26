package main

import (
	"fmt"
	"sync"
)

func main() {
	aa()
}

func aa() {
	defer func() {
		fmt.Print("xxxxioooooo")
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	defer func() {
	}()

	panic("v")
	fmt.Println("ooooo")
	wg.Done()

	wg.Wait()

}

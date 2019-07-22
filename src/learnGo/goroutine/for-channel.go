package main

import "fmt"

func main() {

	numChan := make(chan int, 10)
	i := 0
	for {
		if i > 50 {
			close(numChan)
			break
		}

		i++
		go func(i int) {
			fmt.Println("oooooo----", i)
			numChan <- i
		}(i)
	}
	for num := range numChan {

		fmt.Println("xxxxxx----", num)
	}
}

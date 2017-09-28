package main

import "fmt"

var letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}

func main() {

	chan_n := make(chan bool)
	chan_l := make(chan bool, 1)

	done := make(chan struct{})

	go func() {
		for i := 1; i < 11; i += 2 {
			<-chan_l
			fmt.Print(i)
			fmt.Print(i + 1)
			chan_n <- true
		}
	}()

	go func() {

		for i := 0; i < 10; i += 2 {
			<-chan_n
			fmt.Print(letters[i])
			fmt.Print(letters[i+1])
			chan_l <- true
		}

		done <- struct{}{}
	}()

	chan_l <- true
	<-done
}

package main

import "fmt"

func main() {

	// c := make(chan bool) //阻塞
	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go sum(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}

	fmt.Println("DONE")
}

func sum(c chan bool, n int) {

	x := 0
	for i := 1; i < 1000000000; i++ {
		x += i
	}
	fmt.Println(n, x)
	c <- true
}

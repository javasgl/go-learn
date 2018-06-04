package main

import "fmt"

func main() {
	done := make(chan string, 1)

	done <- "yes"
	close(done) //close 也会导致case <-done 执行

	select {
	case val := <-done:
		fmt.Println("val", val)
		fmt.Println("close......")
	}

}

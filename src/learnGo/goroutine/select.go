package main

import "fmt"

//随机先ch中写入0或者1
func main() {

	count := 0
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 1:
		case ch <- 0:
		}

		i := <-ch

		fmt.Println("ch recevied:", i)
		count++
		if count > 100 {
			fmt.Println(count)
			break
		}
	}

}

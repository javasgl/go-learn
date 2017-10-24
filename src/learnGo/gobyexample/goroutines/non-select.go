package main

import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("recevied msg:", msg)
	default:
		fmt.Println("no msg recevied")
	}

	h1 := "hello"
	select {
	case messages <- h1:
		fmt.Println("send msg:", h1)
	default:
		fmt.Println("no msg send")
	}

	select {
	case msg := <-messages:
		fmt.Println("reveived msg:", msg)
	case sig := <-signals:
		fmt.Println("recevied signal:", sig)
	default:
		fmt.Println("no activity")
	}

}

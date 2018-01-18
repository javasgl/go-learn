package main

import "fmt"

func main() {

	defer func() {

		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("no")
		}
	}()
	defer func() {

		panic("1111")
	}()
	panic("2222222")
}

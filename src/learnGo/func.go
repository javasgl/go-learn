package main

import (
	"fmt"
)

func main() {
	fmt.Println("main....")
	A(2, 4, 6)
}

func A(a ...int) {

	fmt.Println(a)
}

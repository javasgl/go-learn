package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	var res int
	res = max(15, 6)
	fmt.Println("max values is :", res)

	var arg1 int
	arg1 = 10
	switch arg1 {
	case 5:
		fmt.Println("arg1 value is 5")
	case 10:
		fmt.Println("arg1 value is 10")
	case 15:
		fmt.Println("arg1 value is 15")
	default:
		fmt.Println("arg1 value is not 10")
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

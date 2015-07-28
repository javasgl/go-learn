package main

import (
	"fmt"
)

/**
 * 闭包
 */
func main() {
	f1 := closure(10)
	fmt.Println(f1(2))
	fmt.Println(f1(5))
}
func closure(x int) func(y int) int {
	return func(y int) int {
		fmt.Println(y)
		fmt.Println(x)
		return x + y
	}
}

package main

import "fmt"

func main() {

	var slice = []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	fmt.Println(slice[:5])
	// fmt.Println(slice[1:10]) //slice bounds out of range
}

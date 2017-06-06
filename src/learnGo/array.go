package main

import "fmt"

func main() {
	var arr1 []int
	printArr(arr1)
	arr1 = append(arr1, 0)
	printArr(arr1)
	arr2 := make([]int, len(arr1), cap(arr1)*2)
	printArr(arr2)
}
func printArr(arr []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(arr), cap(arr), arr)
}

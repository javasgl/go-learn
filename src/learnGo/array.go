package main

import "fmt"

func main() {

	var array1 [5]*string
	var array2 = [5]*string{new(string)}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array2[0])

	var arr1 []int
	printArr(arr1)
	arr1 = append(arr1, 0)
	printArr(arr1)
	arr2 := make([]int, len(arr1), cap(arr1)*2)
	printArr(arr2)

	slice := make([]int, 5)
	fmt.Println(slice)
	for i := 0; i < 5; i++ {
		slice = append(slice, i) //会累加，导致slice长度发生变化
		slice[i] = i
	}
	fmt.Println(slice)
}
func printArr(arr []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(arr), cap(arr), arr)
}

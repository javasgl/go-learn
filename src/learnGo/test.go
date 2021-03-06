package main

import "fmt"

const (
	SIZE = iota
	SIZE_1
	SIZE_2
	SIZE_3
	SIZE_4
	Font = 200
)
const (
	SIZE_5   = iota
	SIZE_5_1 = 2

	SIZE_6 = iota
)

func main() {

	fmt.Println(Font)
	fmt.Println("size:", SIZE)
	fmt.Println("size1:", SIZE_1)
	fmt.Println("size2:", SIZE_2)
	fmt.Println("size3:", SIZE_3)
	fmt.Println("size4:", SIZE_4)
	fmt.Println("size5:", SIZE_5)
	fmt.Println("size5_1:", SIZE_5_1)
	fmt.Println("size6:", SIZE_6)

	fmt.Println("Hello World")
	arr := []int{1, 3, 4, 5, 6, 7}
	for index, val := range arr {

		fmt.Println(index, "===", val)
	}

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

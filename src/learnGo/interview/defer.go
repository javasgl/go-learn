package main

import "fmt"

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("1.0", a, b))
	a = 0
	defer calc("2", a, calc("2.0", a, b))
	b = 1

}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)

	return ret
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 97
	var b = strconv.Itoa(a)
	fmt.Println(strconv.Itoa(a))
	fmt.Println(strconv.Atoi(b))
}

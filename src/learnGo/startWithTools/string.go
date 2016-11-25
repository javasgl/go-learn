package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "java sgl"
	arr := strings.Split(name, " ")
	fmt.Println(arr)
	arr2 := strings.Join(arr, "*")
	fmt.Println(arr2)
	str := strings.Replace(name, "sgl", "songgl", -1)
	fmt.Println(str)
}

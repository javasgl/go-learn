package main

import (
	"fmt"
	"learnGo/stringutil"
	"strconv"
)

func main() {
	var a int = 97
	var b = strconv.Itoa(a)
	fmt.Println(strconv.Itoa(a))
	fmt.Println(strconv.Atoi(b))
	fmt.Println(stringutil.Say("Hello world"))
}

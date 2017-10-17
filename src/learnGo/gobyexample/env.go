package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("foo", "1")

	fmt.Println("foo:", os.Getenv("foo"))
	fmt.Println("bar:", os.Getenv("bar"))

	for _, e := range os.Environ() {

		// fmt.Println(e)
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], "=", pair[1])
	}

}

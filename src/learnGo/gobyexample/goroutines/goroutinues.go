package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from+":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(from string) {

		fmt.Println(from)

	}("going")

	var input string
	fmt.Scanln(&input)

	for len(input) == 0 {
		fmt.Println("please input something:")
		fmt.Scanln(&input)
	}
	fmt.Println("done:", len(input), input)
}

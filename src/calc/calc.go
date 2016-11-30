package main

import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

var Usage = func() {

	fmt.Println("Usage:calc command [arguments] ...")
	fmt.Println("\nThe commands are :\n\tadd\taddition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main() {
	for i, v := range os.Args {
		fmt.Printf("%d:\t%s\n", i, v)
	}
	args := os.Args[1:]
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	switch args[0] {
	case "add":
		if len(args) != 3 {
			fmt.Println("Usage calc add <integer><interger2>")
			return
		}
		v1, err1 := strconv.Atoi(args[1])
		v2, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			fmt.Println("Usage calc add <integer><interger2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 2 {
			fmt.Println("Usage calc sqrt <integer>")
			return
		}
		v, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Usage calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("ResultL: ", ret)
	default:
		Usage()
	}
}

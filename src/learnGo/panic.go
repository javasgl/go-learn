package main

import (
	"errors"
	"fmt"
)

func innerFunc() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Fatal Error: ", p)
		}
	}()
	fmt.Println("Enter innerFunc...")
	panic(errors.New("Occur a Error!"))
	fmt.Println("Quit innerFunc...")
}
func outerFunc() {
	fmt.Println("Enter outerFunc....")
	innerFunc()
	fmt.Println("Quit outerFunc...")

}
func main() {
	fmt.Println("Enter main....")
	outerFunc()
	fmt.Println("Quit main....")
}

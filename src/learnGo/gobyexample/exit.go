package main

import (
	"fmt"
	"os"
)

//go run exit.go
//exit status 3
//go build exit.go

//./exit
//echo $?
//3
func main() {

	defer fmt.Println("done!")
	os.Exit(3)
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	filename   = "sample.txt"
	start_data = "12345"
)

func printContents() {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func main() {

	// printContents()

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return
	}
	//if _, err := f.Seek(20, 0); err != nil {
	//	panic(err)
	//}

	if _, err := f.WriteAt([]byte("A"), info.Size()); err != nil {
		panic(err)
	}

	printContents()
}

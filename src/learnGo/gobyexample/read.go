package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	code, err := ioutil.ReadFile("signals.go")
	check(err)
	fmt.Println(string(code))
	fmt.Println("------")

	f, err := os.Open("signals.go")
	check(err)
	defer f.Close()
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes:%s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 7)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d:%s\n", n2, o2, string(b2))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

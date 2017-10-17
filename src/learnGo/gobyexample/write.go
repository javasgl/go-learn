package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

func main() {

	d1 := []byte("hello,golang")

	err := ioutil.WriteFile("tmp", d1, 0644)

	check(err)

	f, err := os.Create("tmp")
	check(err)
	defer f.Close()
	_, err = f.WriteString("second,hello,golang")
	check(err)
	f.Sync()

	w := bufio.NewWriter(f)
	_, err = w.WriteString("\nbuffered\n")
	w.Flush()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var init_string string
var once sync.Once

func main() {
	Do()
	time.Sleep(time.Millisecond)
}

func setup() {
	fmt.Println("setup....should be run only once!")
	init_string = "init the global config or data"
}

func Do() {
	go print()
	go print()
}
func print() {
	once.Do(setup)
	// setup()
	fmt.Println(init_string)
}

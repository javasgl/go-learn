package main

import (
	"fmt"
	"plugin"
)

func main() {

	p, err := plugin.Open("p1.so")
	if err != nil {
		panic(err)
	}
	m, err := p.Lookup("GetProductBasePrice")
	if err != nil {
		panic(err)
	}
	res := m.(func() int)()
	fmt.Println(res)

	p2, err := plugin.Open("p2.so")
	if err != nil {
		panic(err)
	}
	m2, err := p2.Lookup("GetProductBasePrice")
	if err != nil {
		panic(err)
	}
}

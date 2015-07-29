package main

import (
	"fmt"
)

type A struct {
	Name string
}
type B struct {
	Name string
}
type C struct {
	Name string
}

func (a A) say() {
	fmt.Println("A_say:" + a.Name)
}
func (b B) say() {
	fmt.Println("B_say:" + b.Name)
}

/**
 * 指针
 */
func (c *C) say() {
	c.Name = "change C_name"
	fmt.Println("c_name::" + c.Name)
}
func main() {
	a := A{Name: "A_name"}
	a.say()
	b := B{Name: "B_name"}
	b.say()
	c := C{Name: "C_name"}
	c.say()
	fmt.Println(c.Name)
}

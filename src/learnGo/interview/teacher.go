package main

import "fmt"

func main() {
	t := Teacher{}
	t.ShowA()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("people`s A")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("people`s B")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher`s B")
}

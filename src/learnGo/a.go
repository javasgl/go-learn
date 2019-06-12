package main

type data struct {
	x int
}

func main() {
	d := data{100}
	var t interface{} = d
	println(t.(data).x)
}

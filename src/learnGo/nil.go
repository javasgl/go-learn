package main

import "fmt"

type People interface {
	Show()
}
type Student struct{}

func (stu *Student) Show() {

}
func live() People {
	var stu *Student
	return stu
}
func live2() People {
	return nil
}

func main() {
	res := live()
	fmt.Printf("%#v", res)
	if res == nil {
		fmt.Println(":res is nil")
	} else {
		fmt.Println(":res not nil")
	}

	res2 := live2()
	fmt.Printf("%#v", res2)
	if res2 == nil {
		fmt.Println(":res2 is nil")
	} else {
		fmt.Println(":res2 not nil")
	}
}

package main

import "fmt"

type Body struct {
	Weight int
}

func main() {
	body := &Body{}
	fmt.Println(body)
	fmt.Println(body.Weight)

	fmt.Println("--------")

	body2 := Body{}
	fmt.Println(body2)
	fmt.Println(body2.Weight)

	fmt.Println("--------")

	b3 := &body2
	fmt.Println(b3)
	fmt.Println(b3.Weight)    //直接使用 p.x 不需要显式 使用 *p.x
	fmt.Println((*b3).Weight) //显式使用 *p.x

	fmt.Println("--------")
	b4 := *b3
	fmt.Println(b4)
	fmt.Println(b4.Weight)
	fmt.Println("--------")

	s := Gen("param1")
	fmt.Println(s)
	fmt.Println(&s)
	fmt.Println(*s)

	fmt.Println("--------")

	fmt.Println(new(Object))
	fmt.Println(&Object{})
	fmt.Println(Object{})

}

type Object map[string][]string

func Gen(param string) (s *Object) {
	return &Object{
		"param":  []string{param},
		"param2": []string{"param2", "tag2"},
		"param3": []string{"param3", "tag3"},
	}
}

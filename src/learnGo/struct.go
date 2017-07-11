package main

import (
	"fmt"
)

type A struct {
	Name string
	B
}
type B struct {
	Name string
}
type C struct {
	Name string
}

type User struct {
	Name string
	Age  int
}

func (user User) Say() {
	fmt.Println("user say:", user.Name, " age:", user.Age)
	user.Age = 1000
	fmt.Println("user say:", user.Name, " age:", user.Age)
}
func (user *User) Say2() {

	fmt.Println("user say2:", user.Name, " age:", user.Age)
	user.Age = 1000
	fmt.Println("user say2:", user.Name, " age:", user.Age)
}

/**
 *  A中包含B
 *  如果A中也包含C，则报错，因为此时B和C是平级关系，会出现两个Name冲突
 *  但是，如果A包含B，B包含C，则不会冲突
 */
func main() {
	a := A{Name: "A_name", B: B{Name: "B_name"}}
	fmt.Println(a.Name, a.B.Name)

	user := User{"songgl", 18}
	user.Say()
	fmt.Println(user)
	(&user).Say()
	fmt.Println(user)
	fmt.Println("=============")

	user.Say2()
	fmt.Println(user)
	(&user).Say2()
	fmt.Println(user)

}

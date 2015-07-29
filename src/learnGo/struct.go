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

/**
 *  A中包含B
 *  如果A中也包含C，则报错，因为此时B和C是平级关系，会出现两个Name冲突
 *  但是，如果A包含B，B包含C，则不会冲突
 */
func main() {
	a := A{Name: "A_name", B: B{Name: "B_name"}}
	fmt.Println(a.Name, a.B.Name)
}

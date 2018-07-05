package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `
package main

import "fmt"

func main() {
	fmt.Println("Hello,AST")
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, f)
}

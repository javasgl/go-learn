package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := fmt.Printf

	po := point{1, 8}

	p("%v\n", po)
	p("%+v\n", po)
	p("%#v\n", po)

	p("%T\n", po)

	p("%t\n", true)
	p("%t\n", false)

	p("%d\n", 12) //十进制
	p("%b\n", 12) //二进制
	p("%c\n", 33) //字符
	p("%x\n", 12) //十六进制

	p("%f\n", 12.458)

	p("%0.2f\n", 12.459)
	p("%8.2f\n", 12.459) //a.b 宽度.精度

	p("%e\n", 123400000.0)
	p("%E\n", 123400000.0)

	p("%q\n", "\"string\"")
	p("%s\n", "\"string\"")

	p("%x\n", "string")

	p("%p\n", &po)

	p("|%6d|%6d|\n", 12, 345)

	p("|%6.2f|%6.2f|\n", 1.2, 3.458) //右对齐

	p("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //左对齐

	p("|%6s|%6s|\n", "foo", "b")
	p("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")

	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

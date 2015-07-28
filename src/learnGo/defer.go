package main

import (
	"fmt"
)

/**
 * defer常用于关闭资源，程序异常处理
 * 先进后出--栈 ->调用顺序
 */
func main() {
	A()
	B()
	C()
}
func A() {
	fmt.Println("A....")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B`s defer")
		}
	}()
	panic("panic in B......")
}
func C() {
	fmt.Println("C.....")
}

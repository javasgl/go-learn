package main

import "fmt"

func main() {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			// return // 虽然能跳出循环，但是也导致程序提前结束，不会执行之后的print out
			break loop //break 到指定tag后不再进入循环,这里不会造成死循环
			//goto loop //goto 到指定tag后会再次进入循环,这里会造成死循环
		}
	}
	fmt.Println("out!")
}

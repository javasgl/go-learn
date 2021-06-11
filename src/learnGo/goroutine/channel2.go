package main

import (
	"fmt"
	"time"
)

// 对已关闭的chan继续读,会返回其内元素类型的默认值
// 对已关闭的channel继续写,会panic
func main() {
	c := make(chan int)

	go func() {
		v, ok := <-c
		fmt.Println(v, ok)
		if !ok {
			// 关闭了，继续读/写
			c <- 1
			for {
				fmt.Print(<-c)
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second)
	close(c)
	time.Sleep(2 * time.Second)
}

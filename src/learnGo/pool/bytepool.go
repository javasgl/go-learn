/**
 * @author songgl
 * @since 2018-06-27 10:21
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	a := time.Now().Unix()
	for i := 0; i <= 1000000000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
	b := time.Now().Unix()
	for i := 0; i <= 1000000000; i++ {
		obj := bytePoll.Get().(*[]byte)
		_ = obj
		bytePoll.Put(obj)
	}
	c := time.Now().Unix()
	fmt.Println("without pool", b-a)
	fmt.Println("with    pool", c-b)
}

var bytePoll = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

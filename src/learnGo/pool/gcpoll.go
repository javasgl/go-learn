/**
 * @author songgl
 * @since 2018-06-27 11:20
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	a := p.Get().(int)
	p.Put(1)
	b := p.Get().(int)
	fmt.Println(a, b)

	a = p.Get().(int)
	p.Put(1)
	runtime.GC() //手动调用GC
	b = p.Get().(int)
	fmt.Println(a, b)
}

package main

import (
	"fmt"
	"reflect"
)

func main() {

	var slice = []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(slice)
	fmt.Println(slice[:5])
	// fmt.Println(slice[1:10]) //slice bounds out of range

	fmt.Println(append(slice, 9, 10))

	var slice2 = []int{7, 8}
	fmt.Println(append(slice, slice2...))

	fmt.Println(reflect.TypeOf(slice).Kind())

	fmt.Println(reflect.TypeOf(make([]interface{}, 0)).Kind())

	var arrSlice = make([]int, 0)

	fmt.Println(arrSlice)
	fmt.Println(len(arrSlice))
	fmt.Println(cap(arrSlice))

	arrSlice = append(arrSlice, 0)

	fmt.Println(arrSlice)
	fmt.Println(len(arrSlice))
	fmt.Println(cap(arrSlice))
	arrSlice = append(arrSlice, 0)

	fmt.Println(arrSlice)
	fmt.Println(len(arrSlice))
	fmt.Println(cap(arrSlice))

	fmt.Println("-----------")

	var as = make([]int, 0, 1)
	as = append(as, 1)
	fmt.Println(as)
	fmt.Println(len(as))
	fmt.Println(cap(as))

	as = append(as, 1)
	fmt.Println(as)
	fmt.Println(len(as))
	fmt.Println(cap(as))

	as = append(as, 1) //扩容*2
	fmt.Println(as)
	fmt.Println(len(as))
	fmt.Println(cap(as))

}

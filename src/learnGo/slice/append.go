package main

import "fmt"

func main() {

	var t = make([]int, 0, 10)
	var s = make([]int, 0, 10)

	fmt.Printf("t:addr:%p\tlen:%v\tcontent:%v\n", t, len(t), t)
	fmt.Printf("s:addr:%p\tlen:%v\tcontent:%v\n", s, len(s), s)

	fmt.Println(".....append....")

	t = append(s, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)

	fmt.Printf("t:addr:%p\tlen:%v\tcontent:%v\n", t, len(t), t)
	fmt.Printf("s:addr:%p\tlen:%v\tcontent:%v\n", s, len(s), s)

	// 切片是基于数组的。可以把切片简单的理解为数组中的一段，
	// 它保留了指向数组中某个元素（可以不是第一个元素）的指针，以及它本身的长度（也就是切片的大小）。
	// var s=make([]int,0,10)这条语句创建了一个长度为10的数组，
	// 并在该数组之上建立了一个切片s，该切片保留指向数组第一个元素的指针， 以及它的长度（0）。
	// t = append(s,1,2,3,4) 这条语句向 s 指向的数组中追加了 4 个元素（注意这里 s 并没有变化，它的长度仍然是0，仅仅是底层数组变了）
	// 并返回新的切片（指向数组的第一个元素，长度为4）赋值给 t 。
	// 现在 s 和 t 都指向同一个数组的第一个元素，因此他们的地址是相同的。而 s 的长度是0，t 的长度是4.
	// 可以参考go的文档：go doc builtin.make 以及 go doc builtin.append

}

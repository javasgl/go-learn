package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

//自定义排序，实现sort.Interface 的三个方法
func main() {

	fruits := []string{"peach", "banana", "pear", "kiwi"}

	sort.Sort(ByLength(fruits))

	fmt.Println(fruits)

}

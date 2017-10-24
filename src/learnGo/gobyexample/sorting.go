package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{"c", "a", "b"}
	sort.Strings(s)
	fmt.Println(s)

	i := []int{4, 2, 6}
	sort.Ints(i)
	fmt.Println(i)

	fmt.Println(sort.IntsAreSorted(i))
}

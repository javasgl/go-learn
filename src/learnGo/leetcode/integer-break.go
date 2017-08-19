package main

import "fmt"

//https://leetcode.com/problems/integer-break/description/
func main() {

	fmt.Println(integerBreak(4))
	fmt.Println(integerBreak(7))
	fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {

	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	max := 1
	for {
		if n > 4 {
			max *= 3
			n -= 3
		} else {
			break
		}
	}
	max *= n

	return max
}

package main

import "fmt"

//https://leetcode.com/problems/counting-bits/description/
func main() {

	fmt.Println(countBits(5))
}

func countBits(num int) []int {

	counts := make([]int, num+1)

	for i := 1; i <= num; i++ {
		counts[i] = counts[i&(i-1)] + 1
	}

	return counts
}

// f(0) = 00000000  0

// f(1) = 00000001  1

// f(2) = 00000010  1

// f(3) = 00000011  2

// f(4) = 00000100  1

// f(5) = 00000101  2

// f(6) = 00000110  2

// f(7) = 00000111  3

// f(8) = 00001000  1

// f(9) = 00001001  2

// f(10) = 00001010  2

// f(11) = 00001011  3

// f(12) = 00001100  2

// f(13) = 00001101  3

// f(14) = 00001110  3

//f(n) = f(n&(n-1)) +1

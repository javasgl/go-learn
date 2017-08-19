package main

import "fmt"

//动态规划--爬楼问题
//https://leetcode.com/problems/climbing-stairs/description/
func main() {

	fmt.Println("3:", climbStairs(3))
	fmt.Println("4:", climbStairs(4))
	fmt.Println("5:", climbStairs(5))
	fmt.Println("6:", climbStairs(6))

	fmt.Println("3:", climbStairs2(3))
	fmt.Println("4:", climbStairs2(4))
	fmt.Println("5:", climbStairs2(5))
	fmt.Println("6:", climbStairs2(6))
}

//递归--超时
func climbStairs(n int) int {

	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

//递推
func climbStairs2(n int) int {

	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return n
	}

	prev_1 := 1
	prev_2 := 2
	next := 0

	for i := 3; i <= n; i++ {
		next = prev_1 + prev_2
		prev_1 = prev_2
		prev_2 = next
	}
	return next
}

package main

import (
	"fmt"
	"math"
)

func main() {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(nums))

	nums = []int{-2, -1}

	fmt.Println(maxSubArray(nums))

}

//dp
//https://leetcode.com/problems/maximum-subarray/description/
//[-2,1,-3,4,-1,2,1,-5,4]
//sum [4,-1,2,1] = 6
func maxSubArray(nums []int) int {

	nums_len := len(nums)

	if nums_len == 0 {
		return 0
	}
	if nums_len == 1 {
		return nums[0]
	}

	global := math.MinInt32
	local := 0

	for _, val := range nums {

		if local+val > val { //means local is postive integer
			local = local + val
		} else {
			local = val
		}

		if local > global { //局部最优 larger then 全局最优
			global = local
		}
	}

	return global
}

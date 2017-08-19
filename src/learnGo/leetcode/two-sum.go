package main

import "fmt"

//主要是array , hash
//https://leetcode.com/problems/two-sum/description/
func main() {

	nums := []int{1, 3, 5, 6, 7, 8, 9}
	target := 9 //1+8 or 3+6

	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {

	numsMap := make(map[int]int, len(nums))

	for index, val := range nums {
		if j, ok := numsMap[target-val]; ok && j != index {
			return []int{j, index}
		}
		numsMap[val] = index
	}
	return nil
}

package main

import "fmt"

//dp
//https://leetcode.com/problems/house-robber/description/
func main() {

	nums := []int{2, 4, 5, 6, 3, 12, 4}

	fmt.Println(rob(nums))
	fmt.Println(rob2(nums))

}

//O(1) space, O(N) time
func rob(nums []int) int {

	houseCount := len(nums)
	if houseCount <= 0 {
		return 0
	}
	if houseCount == 1 {
		return nums[0]
	}

	prevNo := 0
	prevYes := 0
	for _, amount := range nums {
		temp := prevNo
		if prevNo < prevYes {
			prevNo = prevYes
		}
		prevYes = temp + amount
	}
	if prevNo > prevYes {
		return prevNo
	} else {
		return prevYes
	}
}

//O(N) space, O(N) time
func rob2(nums []int) int {
	houseCount := len(nums)
	if houseCount <= 0 {
		return 0
	}
	if houseCount == 1 {
		return nums[0]
	}

	max := make([]int, houseCount)

	max[0] = nums[0]
	max[1] = nums[0] //dont rob,smae with the before

	if nums[1] > nums[0] {
		max[1] = nums[1] //rob if larger then dont
	}

	for i := 2; i < houseCount; i++ {

		max[i] = max[i-1] //dont rob,same with the before

		if nums[i]+max[i-2] > max[i] {

			max[i] = nums[i] + max[i-2] //rob if larger then dont

		}

	}
	return max[houseCount-1]
}

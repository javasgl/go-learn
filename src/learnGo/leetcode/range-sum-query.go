package main

import "fmt"

//https://leetcode.com/problems/range-sum-query-immutable/description/
func main() {

	obj := Constructor([]int{1, 2, 3, 5, 6})

	fmt.Println(obj.SumRange(0, 1))

	fmt.Println(obj.SumRange(0, 2))

	fmt.Println(obj.SumRange(2, 2))

}

type NumArray struct {
	rangeSum []int
}

func Constructor(nums []int) NumArray {

	rangeSum := make([]int, len(nums))

	if len(nums) > 0 {
		rangeSum[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			rangeSum[i] = rangeSum[i-1] + nums[i]
		}
	}
	return NumArray{rangeSum}
}

func (this *NumArray) SumRange(i, j int) int {

	if i == 0 {
		return this.rangeSum[j]
	} else {
		return this.rangeSum[j] - this.rangeSum[i-1]
	}
}

package main

import (
	"fmt"
	"math"
)

//股票最大收益
//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func main() {

	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))

	fmt.Println(maxProfit2(prices))
}

//暴力解法
func maxProfit(prices []int) int {
	maxProfit := 0
	for i, i_price := range prices {
		subPrices := prices[i+1:]
		for _, j_price := range subPrices {
			profit := j_price - i_price
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPirce := math.MaxInt32
	maxProfit := 0
	for _, v := range prices {
		if v-minPirce > maxProfit {
			maxProfit = v - minPirce
		}
		if v < minPirce {
			minPirce = v
		}
	}
	return maxProfit
}

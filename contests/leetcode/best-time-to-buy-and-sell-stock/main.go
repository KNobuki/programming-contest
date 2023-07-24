package main

import "fmt"

func maxProfit(prices []int) int {
	max := make([]int, len(prices))
	max[len(prices)-1] = prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > max[i+1] {
			max[i] = prices[i]
		} else {
			max[i] = max[i+1]
		}
	}
	min := prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		if ans < max[i]-min {
			ans = max[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{1, 2}))
}

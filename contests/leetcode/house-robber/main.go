package main

import "fmt"

func rob(nums []int) int {
	// dp[i][0] i番目の家に泥棒に入らない
	// dp[i][1] i番目の家に泥棒に入る
	dp := make([][]int, len(nums))
	dp[0] = make([]int, 2)
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

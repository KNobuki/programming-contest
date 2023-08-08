package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp1 := make([][]int, len(nums)-1)
	dp2 := make([][]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		dp1[i] = make([]int, 2)
		dp2[i] = make([]int, 2)
	}
	dp1[0][1] = nums[0]
	dp2[0][1] = nums[1]
	for i := 1; i < len(nums)-1; i++ {
		dp1[i][0] = max(dp1[i-1][0], dp1[i-1][1])
		dp1[i][1] = dp1[i-1][0] + nums[i]
		dp2[i][0] = max(dp2[i-1][0], dp2[i-1][1])
		dp2[i][1] = dp2[i-1][0] + nums[i+1]
	}
	return max(dp1[len(nums)-2][0], max(dp1[len(nums)-2][1], max(dp2[len(nums)-2][0], dp2[len(nums)-2][1])))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

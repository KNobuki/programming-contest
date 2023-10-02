package main

func moveZeroes(nums []int) {
	zeroNum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroNum++
		} else {
			nums[i-zeroNum] = nums[i]
		}
	}
	for i := len(nums) - zeroNum; i < len(nums); i++ {
		nums[i] = 0
	}
}

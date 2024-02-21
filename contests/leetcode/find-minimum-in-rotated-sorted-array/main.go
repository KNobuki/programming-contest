package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for r-l > 1 {
		mid := (l + r) / 2
		if nums[mid] > nums[l] {
			l = mid
		} else {
			r = mid
		}
	}
	if nums[0] < nums[r] {
		return nums[0]
	}
	return nums[r]
}

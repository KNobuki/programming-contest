package main

import "fmt"

func main() {
	fmt.Println(search([]int{3, 1}, 3))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for r-l > 1 {
		mid := (l + r) / 2
		if nums[mid] > nums[l] {
			l = mid
		} else {
			r = mid
		}
	}
	var minIndex int
	if nums[0] < nums[r] {
		minIndex = 0
	} else {
		minIndex = r
	}
	l, r = 0, minIndex
	for r-l > 1 {
		mid := (l + r) / 2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid
		}
	}
	if nums[l] == target {
		return l
	}
	l, r = minIndex, len(nums)
	for r-l > 1 {
		mid := (l + r) / 2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

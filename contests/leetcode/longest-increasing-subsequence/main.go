package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS(nums []int) int {
	var l []int
	for i := 0; i < len(nums); i++ {
		pos := sort.Search(len(l), func(j int) bool {
			return l[j] >= nums[i]
		})
		if pos >= len(l) {
			l = append(l, nums[i])
		} else {
			l[pos] = nums[i]
		}
	}
	return len(l)
}

func main() {
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
}

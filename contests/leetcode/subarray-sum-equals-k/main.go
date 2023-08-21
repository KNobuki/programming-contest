package main

import "fmt"

func subarraySum(nums []int, k int) int {
	ans := 0
	cum := make([]int, len(nums)+1)
	cum[0] = 0
	for i := 0; i < len(nums); i++ {
		cum[i+1] = cum[i] + nums[i]
	}
	for i := 0; i < len(nums)+1; i++ {
		for j := i + 1; j < len(nums)+1; j++ {
			if cum[j]-cum[i] == k {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}

package main

import "math/bits"

func countBits(n int) []int {
	ans := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		ans = append(ans, bits.OnesCount(uint(i)))
	}
	return ans
}

package main

func subsets(nums []int) [][]int {
	ret := make([][]int, 1<<len(nums))
	for bit := 0; bit < 1<<len(nums); bit++ {
		ret[bit] = []int{}
		for i := 0; i < len(nums); i++ {
			if bit&(1<<i) > 0 {
				ret[bit] = append(ret[bit], nums[i])
			}
		}
	}
	return ret
}

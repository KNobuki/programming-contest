package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int
	var dfs func(now []int, remaining, index int)
	dfs = func(now []int, remaining, index int) {
		if remaining == 0 {
			ret = append(ret, now)
			return
		}
		if remaining < 0 {
			return
		}
		for i := index; i < len(candidates); i++ {
			next := make([]int, len(now)+1)
			copy(next, append(now, candidates[i]))
			dfs(next, remaining-candidates[i], i)
		}
	}
	dfs([]int{}, target, 0)
	return ret
}

func main() {
	ret := combinationSum([]int{2, 3, 5}, 8)
	for _, v := range ret {
		for _, vv := range v {
			fmt.Printf("%d ", vv)
		}
		fmt.Println()
	}
}

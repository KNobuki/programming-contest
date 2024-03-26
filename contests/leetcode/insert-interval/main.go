package main

func insert(intervals [][]int, newInterval []int) [][]int {
	ret := make([][]int, 0, len(intervals)+1)
	added := false
	l, r := newInterval[0], newInterval[1]
	for _, interval := range intervals {
		if !added && newInterval[1] < interval[0] {
			added = true
			ret = append(ret, []int{l, r})
		}

		if between(interval[0], newInterval[0], newInterval[1]) || between(interval[1], newInterval[0], newInterval[1]) {
			l = min(l, interval[0])
			r = max(r, interval[1])
		} else if interval[0] <= newInterval[0] && newInterval[1] <= interval[1] {
			added = true
			ret = append(ret, []int{interval[0], interval[1]})
		} else {
			ret = append(ret, []int{interval[0], interval[1]})o
		}
	}
	if !added {
		ret = append(ret, []int{l, r})
	}
	return ret
}

func between(n, a, b int) bool {
	return n >= a && n <= b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

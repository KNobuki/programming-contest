package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcda"))
}

func lengthOfLongestSubstring(s string) int {
	letters := make(map[byte]int)
	que := make([]byte, 0)
	ans := 0
	for _, c := range s {
		que = append(que, byte(c))
		letters[byte(c)]++
		for len(que) > 0 && letters[byte(c)] > 1 {
			head := que[0]
			que = que[1:]
			letters[head]--
		}
		if len(que) > ans {
			ans = len(que)
		}
	}
	return ans
}

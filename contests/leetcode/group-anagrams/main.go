package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, str := range strs {
		sortedStr := []byte(str)
		sort.Slice(sortedStr, func(i, j int) bool {
			return sortedStr[i] < sortedStr[j]
		})
		if _, ok := mp[string(sortedStr)]; !ok {
			mp[string(sortedStr)] = []string{}
		}
		mp[string(sortedStr)] = append(mp[string(sortedStr)], str)
	}
	ret := [][]string{}
	for _, strings := range mp {
		ret = append(ret, strings)
	}
	return ret
}

package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, str := range strs {
		bstr := []byte(str)
		sort.Slice(bstr, func(i, j int) bool {
			return bstr[i] < bstr[j]
		})
		sortedStr := string(bstr)
		if _, ok := mp[sortedStr]; !ok {
			mp[sortedStr] = []string{}
		}
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	var ret [][]string
	for _, strings := range mp {
		ret = append(ret, strings)
	}
	return ret
}

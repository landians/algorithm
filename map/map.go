package _map

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		// 排序后，字符串的判断就很简单了
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}

	anagrams := make([][]string, 0, len(m))
	for _, ss := range m {
		anagrams = append(anagrams, ss)
	}
	return anagrams
}

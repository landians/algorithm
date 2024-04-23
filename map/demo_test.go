package _map

import "testing"

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagrams := groupAnagrams(strs)
	t.Log(anagrams)
}

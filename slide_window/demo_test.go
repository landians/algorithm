package slide_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	l := lengthOfLongestSubstring(s)
	assert.Equal(t, 3, l)

	s = "aaaaaaa"
	l = lengthOfLongestSubstring(s)
	assert.Equal(t, 1, l)

	s = "pwwkew"
	l = lengthOfLongestSubstring(s)
	assert.Equal(t, 3, l)
}

func TestMinWindow(t *testing.T) {
	ss, s := "ADBECFEBANC", "ABC"
	assert.Equal(t, "BANC", minWindow(ss, s))
}

func TestFindAnagrams(t *testing.T) {
	s, p := "cbaebabacd", "abc"
	anagrams := findAnagrams(s, p)
	assert.Equal(t, []int{0, 6}, anagrams)

	s, p = "abab", "ab"
	anagrams = findAnagrams(s, p)
	assert.Equal(t, []int{0, 1, 2}, anagrams)
}

func TestCheckInclusion(t *testing.T) {
	s1, s2 := "oow", "helloworld"
	inclusion := checkInclusion(s1, s2)
	assert.Equal(t, true, inclusion)

	s1, s2 = "ab", "eidbaooo"
	inclusion = checkInclusion(s1, s2)
	assert.Equal(t, true, inclusion)

	s1, s2 = "ab", "eidboaoo"
	inclusion = checkInclusion(s1, s2)
	assert.Equal(t, false, inclusion)
}

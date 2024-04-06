package _34_isPalindrome

import (
	"github.com/landians/algorithm/leetcode"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome_1(t *testing.T) {
	n1 := &leetcode.ListNode{Val: 1}
	n2 := &leetcode.ListNode{Val: 2}
	n3 := &leetcode.ListNode{Val: 3}
	n4 := &leetcode.ListNode{Val: 2}
	n5 := &leetcode.ListNode{Val: 1}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	assert.Equal(t, true, isPalindrome(n1))
}

func TestIsPalindrome_2(t *testing.T) {
	n1 := &leetcode.ListNode{Val: 1}
	n2 := &leetcode.ListNode{Val: 2}
	n3 := &leetcode.ListNode{Val: 2}
	n4 := &leetcode.ListNode{Val: 1}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	assert.Equal(t, true, isPalindrome(n1))
}

func TestIsPalindrome_3(t *testing.T) {
	n1 := &leetcode.ListNode{Val: 1}
	n2 := &leetcode.ListNode{Val: 2}
	n3 := &leetcode.ListNode{Val: 3}
	n4 := &leetcode.ListNode{Val: 1}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	assert.Equal(t, false, isPalindrome(n1))
}

func TestIsPalindrome_4(t *testing.T) {
	n1 := &leetcode.ListNode{Val: 1}
	assert.Equal(t, true, isPalindrome(n1))
}

func TestIsPalindrome_5(t *testing.T) {
	n1 := &leetcode.ListNode{Val: 1}
	n2 := &leetcode.ListNode{Val: 1}
	n1.Next = n2
	assert.Equal(t, true, isPalindrome(n1))
}

package _34_isPalindrome

import (
	"github.com/landians/algorithm/leetcode"
	"github.com/landians/algorithm/stack"
)

func isPalindrome(head *leetcode.ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head.Next, head
	n := 0

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		n++
	}

	s := stack.New(n + 1)
	for slow != nil {
		s.Push(slow.Val)
		slow = slow.Next
	}

	p := head
	for !s.IsEmpty() {
		nv := s.Pop().(int)
		if nv != p.Val {
			return false
		}
		p = p.Next
	}

	return true
}

package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func FormatChain(head *ListNode) string {
	if head == nil {
		return ""
	}

	format := "head"

	p := head
	for p != nil {
		format += fmt.Sprintf("->%d", p.Val)
		p = p.Next
	}

	return format
}

func NewListFromArray(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: -999}
	p := head
	for _, value := range values {
		p.Next = &ListNode{Val: value}
		p = p.Next
	}

	return head.Next
}

package _6_partition

import (
	"github.com/landians/algorithm/leetcode"
)

func partition(head *leetcode.ListNode, x int) *leetcode.ListNode {
	if head == nil {
		return nil
	}

	var (
		lh, lt *leetcode.ListNode // 记录小于 x 的节点
		hh, ht *leetcode.ListNode // 记录大于/等于 x 的节点
	)

	p := head
	for p != nil {
		// 断开当前节点
		next := p.Next
		p.Next = nil

		if p.Val < x {
			if lh == nil && lt == nil {
				lh = p
				lt = p
			} else {
				lt.Next = p
				lt = lt.Next
			}
		} else {
			if hh == nil && ht == nil {
				hh = p
				ht = p
			} else {
				ht.Next = p
				ht = ht.Next
			}
		}

		p = next
	}

	if lt != nil {
		lt.Next = hh
	}

	if lh != nil {
		return lh
	}

	return head
}

package xx_partition

import "github.com/landians/algorithm/leetcode"

func partition(head *leetcode.ListNode, x int) *leetcode.ListNode {
	if head == nil {
		return nil
	}

	var (
		lt, lh *leetcode.ListNode // [lt, lh] 是小于 x 的区间
		et, eh *leetcode.ListNode // [et, eh] 是等于 x 的区间
		ht, hh *leetcode.ListNode // [ht, hh] 是大于 x 的区间
	)

	p := head
	for p != nil {
		// 断开当前节点
		next := p.Next
		p.Next = nil

		switch {
		case p.Val < x:
			if lt == nil && lh == nil {
				lt = p
				lh = p
			} else {
				lt.Next = p
				lt = lt.Next
			}
		case p.Val == x:
			if et == nil && eh == nil {
				et = p
				eh = p
			} else {
				et.Next = p
				et = et.Next
			}
		case p.Val > x:
			if ht == nil && hh == nil {
				ht = p
				hh = p
			} else {
				ht.Next = p
				ht = ht.Next
			}
		}
		p = next
	}

	// 将[lh, lt], [eh, et], [mh, mt] 三个区间串联起来, 需要考虑 [lh, lt], [eh, et] 这两个区域可能不存在的情况
	if lt != nil {
		lt.Next = eh
	}

	if et != nil {
		et.Next = hh
	}

	if lh != nil {
		return lh
	}

	return head
}

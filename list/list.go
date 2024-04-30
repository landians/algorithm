package list

import (
	"fmt"
)

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

func reverseLinkedList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// https://leetcode.cn/problems/reverse-linked-list-ii/solutions/634701/fan-zhuan-lian-biao-ii-by-leetcode-solut-teyq/
func reverseBetween(head *ListNode, left, right int) *ListNode {
	// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// 第 1 步：从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
	// 建议写在 for 循环里，语义清晰
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 第 2 步：从 pre 再走 right - left + 1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 第 3 步：切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	// 注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 第 4 步：同第 206 题，反转链表的子区间
	_ = reverseLinkedList(leftNode)

	// 第 5 步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var (
		lh, lt *ListNode // 记录小于 x 的节点
		hh, ht *ListNode // 记录大于/等于 x 的节点
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

func partitionx(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var (
		lt, lh *ListNode // [lt, lh] 是小于 x 的区间
		et, eh *ListNode // [et, eh] 是等于 x 的区间
		ht, hh *ListNode // [ht, hh] 是大于 x 的区间
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

// https://leetcode.cn/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	dummy := &ListNode{Val: -1}
	p := dummy
	p1, p2 := list1, list2

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}

		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}

// https://leetcode.cn/problems/merge-k-sorted-lists/description/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummy := &ListNode{Val: -1}
	p := dummy

	heap := NewHeap(len(lists))

	// 将 k 个链表的头结点加入最小堆
	for _, head := range lists {
		if head != nil {
			heap.Push(head)
		}
	}

	for heap.Len() > 0 {
		// 获取最小节点，接到结果链表中
		node := heap.Pop()
		p.Next = node

		// 将下一个节点加入到堆中
		if node.Next != nil {
			heap.Push(node.Next)
		}

		p = p.Next
	}

	return dummy.Next
}

// https://leetcode.cn/problems/middle-of-the-linked-list/description/
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 快慢指针初始化指向 head
	slow, fast := head, head

	// 快指针走到末尾时停止
	for fast != nil && fast.Next != nil {
		// 慢指针走一步，快指针走两步
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 慢指针指向中点
	return slow
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// 快慢指针初始化指向 head
	slow, fast := head, head

	// 快指针走到末尾时停止
	for fast != nil && fast.Next != nil {
		// 慢指针走一步，快指针走两步
		slow = slow.Next
		fast = fast.Next.Next

		// 快慢指针相遇，说明含有环
		if slow == fast {
			return true
		}
	}

	// 不包含环
	return false
}

// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// p1 指向 A 链表头结点，p2 指向 B 链表头结点
	p1, p2 := headA, headB

	for p1 != p2 {
		// p1 走一步，如果走到 A 链表末尾，转到 B 链表
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		// p2 走一步，如果走到 B 链表末尾，转到 A 链表
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}

	return p1
}

// 主函数
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 虚拟头结点
	dummy := &ListNode{-1, head}

	// 删除倒数第 n 个，要先找倒数第 n + 1 个节点
	x := findFromEnd(dummy, n+1)

	// 删掉倒数第 n 个节点
	x.Next = x.Next.Next

	return dummy.Next
}

// 返回链表的倒数第 k 个节点
func findFromEnd(head *ListNode, k int) *ListNode {
	p1 := head

	// p1 先走 k 步
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}

	p2 := head
	// p1 和 p2 同时走 n - k 步
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	// p2 现在指向第 n - k + 1 个节点，即倒数第 k 个节点
	return p2
}

// https://leetcode.cn/problems/palindrome-linked-list/description/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var pre *ListNode

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 此时 slow 指向了中间节点, 先断开后半部分节点
	pre.Next = nil

	// 将后半部分链表反转
	re := reverseLinkedList(slow)
	p1 := head
	p2 := re

	// 开始判断是否为回文链表
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	se := reverseLinkedList(re)
	pre.Next = se

	return false
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}

	// 断开与后面重复元素的连接
	slow.Next = nil
	return head
}

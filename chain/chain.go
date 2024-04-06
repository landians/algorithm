package chain

import (
	"container/heap"
	"fmt"
	"github.com/landians/algorithm/stack"
	"math/rand"
	"time"
)

const (
	defaultSize = 10
)

// Node 单链表的节点
type Node struct {
	value int
	next  *Node
}

func newNode(v int) *Node {
	return &Node{value: v}
}

// Chain 单链表
type Chain struct {
	head *Node
	len  int
}

func newChain() *Chain {
	return &Chain{head: newNode(0)}
}

func (ch *Chain) insert(v int) *Node {
	vNode := newNode(v)
	cur := ch.head

	for cur.next != nil {
		cur = cur.next
	}

	cur.next = vNode

	ch.len++

	return vNode
}

func (ch *Chain) delete(no *Node) *Node {
	if no == nil {
		return nil
	}

	pre := ch.head
	cur := ch.head.next

	for cur != nil {
		if cur == no {
			break
		}

		pre = cur
		cur = cur.next
	}

	// 找不到待删除的节点 v
	if cur == nil {
		return nil
	}

	// 执行删除
	pre.next = cur.next
	cur = nil

	ch.len--
	return cur
}

func (ch *Chain) query(v int) *Node {
	cur := ch.head

	for cur != nil {
		if cur.value == v {
			break
		}

		cur = cur.next
	}

	return cur
}

func (ch *Chain) update(no *Node, v int) *Node {
	if no == nil {
		return nil
	}

	cur := ch.head

	for cur != nil {
		if cur == no {
			break
		}

		cur = cur.next
	}

	if cur != nil {
		cur.value = v
	}
	return cur
}

func (ch *Chain) traverse() {
	if ch.head == nil || ch.head.next == nil {
		return
	}

	format := "head"

	cur := ch.head.next

	for cur != nil {
		format += fmt.Sprintf("->%d", cur.value)
		cur = cur.next
	}

	fmt.Println(format)
}

// IsPalindrome1 用于判断单链表是否是回环结构， 使用栈
func IsPalindrome1(head *Node) bool {
	if head == nil {
		return true
	}

	// 准备一个大小和链表长度一样的栈
	s := stack.New(defaultSize)

	// 遍历链表，将所有的节点压入栈中
	cur := head
	for cur != nil {
		s.Push(cur)
		cur = cur.next
	}

	// 重新遍历链表，并从栈中依次弹出节点进行比对，若有一个节点不相同，则返回 false
	cur = head
	for cur != nil {
		vNode := s.Pop().(*Node)
		if vNode.value != cur.value {
			break
		}
		cur = cur.next
	}

	return cur == nil
}

// IsPalindrome2 用于判断单链表是否是回环结构， 使用栈
func IsPalindrome2(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}

	// 使用快慢指针的方式
	fast := head
	slow := head.next

	// 快指针一次走两步，慢指针一次走一步，当快指针走完时，慢指针恰好走到了链表的中间
	// 若链表的长度为奇数，那么最后慢指针指向的是 链表的长度 / 2 + 1 的位置，
	// 若链表的长度为偶数，那么最后慢指针指向的是 链表长度的 / 2 的位置
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// 准备一个大小和 链表长度/2 一样的栈
	s := stack.New(defaultSize)

	// 将慢指针指向的节点的后面的节点入栈
	for slow != nil {
		s.Push(slow)
		slow = slow.next
	}

	// 从栈中依次弹出节点，并与链表的前一半节点进行比较, 若有一个节点不相同，则返回 false
	cur := head
	for !s.IsEmpty() {
		vNode := s.Pop().(*Node)
		if cur.value != vNode.value {
			return false
		}
		cur = cur.next
	}
	return true
}

// IsPalindrome3 用于判断单链表是否是回环结构， 不使用栈
func IsPalindrome3(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}

	// 使用快慢指针的方式
	fast := head
	slow := head

	// 快指针一次走两步，慢指针一次走一步，当快指针走完时，慢指针恰好走到了链表的中间
	// 若链表的长度为奇数，那么最后慢指针指向的是 链表的长度 / 2 的位置，
	// 若链表的长度为偶数，那么最后慢指针指向的是 链表长度的 / 2 - 1 的位置
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// 将慢指针指向的节点后面的节点逆序
	fast = slow.next // 记录链表中间节点的下一个节点
	slow.next = nil  // 链表中间节点指向nil

	var p *Node

	for fast != nil {
		p = fast.next    // 记录快指针的后驱节点
		fast.next = slow // 逆序操作
		slow = fast      // 慢指针后移
		fast = p         // 快指针后移
	}

	// 记录最后一个节点
	p = slow
	fast = head

	palindrome := true

	// 分别从头节点和尾节点开始比较数据，若有一个不同，则返回false
	for fast != nil && slow != nil {
		if fast.value != slow.value {
			palindrome = false
			break
		}

		fast = fast.next
		slow = slow.next
	}

	// 开始将之前逆序的链表调整回来
	slow = p.next // 慢指针指向倒数第二个节点
	p.next = nil  // 尾节点的 next 重新指向 nil

	for slow != nil {
		fast = slow.next
		slow.next = p
		p = slow
		slow = fast
	}
	return palindrome
}

// ChainPartition 将单链表按照节点值 < value | = value | > value 进行区域划分
func ChainPartition(head *Node, value int) *Node {
	if head == nil {
		return nil
	}
	// < value 的区间: [lh lt]
	// = value 的区间: [eh et]
	// > value 的区间: [mh mt]
	var (
		lh *Node // less than head
		lt *Node // less than tail
		eh *Node // equal than head
		et *Node // equal than tail
		mh *Node // more than head
		mt *Node // more than tail
	)

	for head != nil {
		// 断开当前节点
		next := head.next
		head.next = nil

		if head.value < value { // 当前节点的值 小于 value，因此放入到 [lh, lt] 这个区间
			if lh == nil {
				lh = head
				lt = head
			} else {
				lt.next = head
				lt = head
			}
		} else if head.value == value { // 当前节点的值 等于 value，因此放入到 [eh, et] 这个区间
			if eh == nil {
				eh = head
				et = head
			} else {
				et.next = head
				et = head
			}
		} else { // 当前节点的值 大于 value，因此放入到 [mh, mt] 这个区间
			if mh == nil {
				mh = head
				mt = head
			} else {
				mt.next = head
				mt = head
			}
		}
		head = next
	}

	// 将[lh, lt], [eh, et], [mh, mt] 三个区间串联起来, 需要考虑 [lh, lt], [eh, et] 这两个区域可能不存在的情况
	if lt != nil {
		lt.next = eh
		if et == nil {
			et = lt
		}
	}

	if et != nil {
		et.next = mh
	}

	if lh != nil {
		return lh
	}

	if eh != nil {
		return eh
	}
	return mh
}

type randNode struct {
	value int
	next  *randNode
	rand  *randNode // 指向了随机的一个节点
}

func newRandNode(value int) *randNode {
	return &randNode{value: value}
}

// CopyListWithRand1 复制含有随机节点指针的单链表，返回复制后的单链表的头节点，前提是原单链表无环， 使用哈希表
func CopyListWithRand1(head *randNode) *randNode {
	// 使用一个哈希表来映射原始节点和复制后的节点的关系
	nodeMap := make(map[*randNode]*randNode)

	// 填充哈希表
	cur := head
	for cur != nil {
		nodeMap[cur] = newRandNode(cur.value)
		cur = cur.next
	}

	// 从头开始设置复制后的单链表的指向
	cur = head
	for cur != nil {
		nodeMap[cur].next = nodeMap[cur.next]
		nodeMap[cur].rand = nodeMap[cur.rand]
		cur = cur.next
	}
	return nodeMap[head]
}

// CopyListWithRand2 复制含有随机节点指针的单链表，返回复制后的单链表的头节点，前提是原单链表无环， 不使用哈希表
func CopyListWithRand2(head *randNode) *randNode {
	if head == nil {
		return nil
	}

	// 将复制的节点插入到每一个原节点的后面，比如: 原来：1->2->3->4 后续: 1->1'->2->2'->3->3'->4->4'
	cur := head
	for cur != nil {
		nextNode := cur.next
		cur.next = newRandNode(cur.value)
		cur.next.next = nextNode
		cur = nextNode
	}

	// 从头节点开始设置每一个复制节点的rand指针指向
	cur = head
	for cur != nil {
		nextNode := cur.next.next
		curCopy := cur.next
		if cur.rand != nil {
			curCopy.rand = cur.rand.next
		} else {
			curCopy.rand = nil
		}
		cur = nextNode
	}

	// 将复制的节点都分离出来组成新的复制后的单链表
	cur = head
	headCopy := head.next
	for cur != nil {
		nextNode := cur.next.next
		curCopy := cur.next
		cur.next = nextNode
		if nextNode != nil {
			curCopy.next = nextNode.next
		} else {
			curCopy.next = nil
		}
		cur = nextNode
	}

	return headCopy
}

func PrintRandChain(head *randNode) {
	if head == nil {
		return
	}

	format := "head"

	cur := head
	for cur != nil {
		format += fmt.Sprintf("->%d", cur.value)
		if cur.rand != nil {
			format += fmt.Sprintf("[rand]%d", cur.rand.value)
		}
		cur = cur.next
	}

	fmt.Println(format)
}

// GetIntersectNode 判断两个单链表是否相交
func GetIntersectNode(head1 *Node, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}

	loop1 := GetLoopNode(head1)
	loop2 := GetLoopNode(head2)

	// 1. 若两个单链表中都有环
	if loop1 != nil && loop2 != nil {
		return bothLoop(head1, head2, loop1, loop2)
	}

	// 2. 若两个单链表都无环，则有可能相交
	if loop1 == nil && loop2 == nil {
		return noLoop(head1, head2)
	}

	// 3.  若两个单链表中只有其中一个有环，则必不可能相交
	return nil
}

// 两个单链表都有环的时候，获取相交的节点
func bothLoop(head1 *Node, head2 *Node, loop1 *Node, loop2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}

	var (
		cur1 *Node
		cur2 *Node
	)

	// 两个单链表的入环节点相同时，以入环节点为终点，重新进行两个无环节点判断相交的过程
	if loop1 == loop2 {
		cur1 = head1
		cur2 = head2

		// n 是用于统计两个单链表长度的差值
		n := 0

		for cur1.next != loop1 {
			n++
			cur1 = cur1.next
		}

		for cur2.next != loop1 {
			n--
			cur2 = cur2.next
		}

		// 若两个单链表无环并且相交的话，则它们的最后一个节点必定相同
		if cur1 != cur2 {
			return nil
		}

		// n > 0 说明 head1 所在的单链表更长, 否则 head2 所在的单链表更长
		if n > 0 {
			cur1 = head1
			cur2 = head2
		} else {
			cur1 = head2
			cur2 = head1
			n = -n
		}

		// 较长的那个单链表从头节点开始走差值步
		for n != 0 {
			n--
			cur1 = cur1.next
		}

		// 接下来两个指针一起走，最后肯定会在相交点出相遇
		for cur1 != cur2 {
			cur1 = cur1.next
			cur2 = cur2.next
		}
		return cur1
	} else {
		// loop1 继续往下走，如果能够走回到自己，则说明两个有环单链表是不相交的
		// 若走的过程中能够碰到loop2, 则说明两个有环单链表是相交的
		cur1 = loop1.next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.next
		}
		return nil
	}
}

// 两个单链表都无环的时候，获取相交的节点
func noLoop(head1 *Node, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}

	cur1 := head1
	cur2 := head2

	// n 是用于统计两个单链表长度的差值
	n := 0

	for cur1.next != nil {
		n++
		cur1 = cur1.next
	}

	for cur2.next != nil {
		n--
		cur2 = cur2.next
	}

	// 若两个单链表无环并且相交的话，则它们的最后一个节点必定相同
	if cur1 != cur2 {
		return nil
	}

	// n > 0 说明 head1 所在的单链表更长, 否则 head2 所在的单链表更长
	if n > 0 {
		cur1 = head1
		cur2 = head2
	} else {
		cur1 = head2
		cur2 = head1
		n = -n
	}

	// 较长的那个单链表从头节点开始走差值步
	for n != 0 {
		n--
		cur1 = cur1.next
	}

	// 接下来两个指针一起走，最后肯定会在相交点出相遇
	for cur1 != cur2 {
		cur1 = cur1.next
		cur2 = cur2.next
	}
	return cur1
}

// GetLoopNode 判断单链表是有环, 返回第一个入环的节点, 不使用哈希表
func GetLoopNode(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}

	// 使用快慢指针的方式
	fast := head
	slow := head

	// 快指针一次走两步，慢指针一次走一步，如果链表有环的话，快指针肯定会再次追上慢指针的
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next

		// 这里就说明快指针再次追上慢指针了
		if fast == slow {
			break
		}
	}

	// 说明链表无环
	if fast.next == nil || fast.next.next == nil {
		return nil
	}

	// 快指针回到头节点
	fast = head

	// 快指针和慢指针一起每次走一步，他们会在入环节点出相遇
	for fast != slow {
		fast = fast.next
		slow = slow.next
	}

	return slow
}

// GetLoopNode2 判断单链表是有环, 返回第一个入环的节点, 使用哈希表
func GetLoopNode2(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodeMap := make(map[*Node]struct{})

	// 遍历链表，并使用哈希表来记录每一个节点，如果链表有环，那肯定遍历的过程中就能够从哈希表中查找到入环节点的值
	cur := head
	for cur != nil {
		if _, ok := nodeMap[cur]; !ok {
			nodeMap[cur] = struct{}{}
		} else {
			break
		}
		cur = cur.next
	}

	return cur
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// AddTwoNumbers 两个链表相加， 复用 l1 或者 l2
func AddTwoNumbers(l1 *Node, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	cur1 := l1
	cur2 := l2
	n := 0

	for cur1 != nil {
		n++
		cur1 = cur1.next
	}

	for cur2 != nil {
		n--
		cur2 = cur2.next
	}

	if n > 0 {
		cur1 = l1
		cur2 = l2
	} else {
		cur1 = l2
		cur2 = l1
	}

	pre, root := cur1, cur1
	addIn := 0

	for cur1 != nil && cur2 != nil {
		twoSum := cur1.value + cur2.value + addIn
		addIn = twoSum / 10
		cur1.value = twoSum % 10
		pre = cur1
		cur1 = cur1.next
		cur2 = cur2.next
	}

	for cur1 != nil {
		twoSum := cur1.value + addIn
		addIn = twoSum / 10
		cur1.value = twoSum % 10
		pre = cur1
		cur1 = cur1.next
	}

	if addIn > 0 {
		pre.next = newNode(addIn)
	}

	return root
}

// AddTwoNumbers2 两个链表相加， 返回新的链表
func AddTwoNumbers2(l1 *Node, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	cur1 := l1
	cur2 := l2
	addIn := 0
	head := newNode(0)
	cursor := head

	for cur1 != nil || cur2 != nil || addIn != 0 {
		cur1Value := 0
		if cur1 != nil {
			cur1Value = cur1.value
		}

		cur2Value := 0
		if cur2 != nil {
			cur2Value = cur2.value
		}

		sumValue := cur1Value + cur2Value + addIn
		addIn = sumValue / 10
		val := sumValue % 10
		cursor.next = newNode(val)
		cursor = cursor.next

		if cur1 != nil {
			cur1 = cur1.next
		}
		if cur2 != nil {
			cur2 = cur2.next
		}
	}

	return head.next
}

func generateRandChain(maxValue int) (*Chain, int) {

	value := rand.Intn(maxValue)

	saveValue := value

	ch := generateChain(saveValue)

	return ch, saveValue
}

func generateFixedChain(value int) (*Chain, int) {
	ch := generateChain(value)

	return ch, value
}

func generateChain(value int) *Chain {
	ch := newChain()
	for value > 0 {
		v := value % 10
		value /= 10
		ch.insert(v)
	}
	return ch
}

func compareChain(ch1 *Chain, ch2 *Chain) bool {
	if ch1 == nil || ch2 == nil {
		return false
	}

	cur1 := ch1.head.next
	cur2 := ch2.head.next

	for cur1 != nil && cur2 != nil {
		if cur1.value != cur2.value {
			break
		}
		cur1 = cur1.next
		cur2 = cur2.next
	}

	return cur1 == nil
}

func ReverseChain(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}

	pre := head
	cur := head.next
	pre.next = nil

	for cur != nil {
		nextNode := cur.next
		cur.next = pre
		pre = cur
		cur = nextNode
	}

	return pre
}

type NodeHeap []*Node

func (h *NodeHeap) Len() int {
	return len(*h)
}

func (h *NodeHeap) Less(i, j int) bool {
	return (*h)[i].value < (*h)[j].value
}

func (h *NodeHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() (v interface{}) {
	// 这里的操作可能是因为内存是按照小端的方式进行存储的
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h *NodeHeap) IsEmpty() bool {
	return h.Len() == 0
}

// KSortChainMerge 将 k 个有序的单链表合并成一个有序链表
func KSortChainMerge(head []*Node) *Node {
	if len(head) == 0 {
		return nil
	}

	root := newNode(0)

	rootNode := root

	nodeHeap := NodeHeap(make([]*Node, 0, len(head)))

	// 将每个链表的头节入堆
	for _, cur := range head {
		nodeHeap = append(nodeHeap, cur)
	}

	for !nodeHeap.IsEmpty() {
		// 从堆中弹出的节点就是最小节点
		cur := heap.Pop(&nodeHeap).(*Node)
		rootNode.next = cur
		rootNode = rootNode.next
		// 依次将弹出节点的下一个节点入堆，本质上是将单链表的后续节点都入堆，然后弹出最小节点
		if cur.next != nil {
			heap.Push(&nodeHeap, cur.next)
		}
	}

	return root.next
}

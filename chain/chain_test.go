package chain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initChain() *Chain {
	ch := newChain()

	for i := 0; i < 10; i++ {
		ch.insert(i)
	}

	return ch
}

func TestChainTraverse(t *testing.T) {
	ch := initChain()
	fmt.Println("===== 单链表遍历 =====")
	ch.traverse()
	fmt.Println("===== 单链表遍历 =====")
}

func TestChainQuery(t *testing.T) {
	ch := initChain()
	fmt.Println("===== 单链表查询 =====")
	no := ch.query(2)
	fmt.Println("query result: ", no.value)
	fmt.Println("===== 单链表查询 =====")
}

func TestChainUpdate(t *testing.T) {
	ch := initChain()
	fmt.Println("===== 单链表更新 =====")
	no := ch.query(2)
	fmt.Println("更新前: ", no.value)
	upNo := ch.update(no, 10)
	fmt.Println("更新后: ", upNo.value)
	fmt.Println("===== 单链表更新 =====")
}

func TestChainDelete(t *testing.T) {
	ch := initChain()
	fmt.Println("===== 单链表删除 =====")
	no := ch.query(2)
	fmt.Println("删除前: ", no)
	delNo := ch.delete(no)
	fmt.Println("删除后: ", delNo)
	fmt.Println("===== 单链表删除 =====")
}

func TestIsPalindrome1(t *testing.T) {
	// 1->2->3->2->1
	ch1 := newChain()
	ch1.insert(1)
	ch1.insert(2)
	ch1.insert(3)
	ch1.insert(2)
	ch1.insert(1)
	assert.Equal(t, IsPalindrome1(ch1.head.next), true)

	// 1->2->2->1
	ch2 := newChain()
	ch2.insert(1)
	ch2.insert(2)
	ch2.insert(2)
	ch2.insert(1)
	assert.Equal(t, IsPalindrome1(ch2.head.next), true)

	// 1->2->3->1->2
	ch3 := newChain()
	ch3.insert(1)
	ch3.insert(2)
	ch3.insert(3)
	ch3.insert(1)
	ch3.insert(2)
	assert.Equal(t, IsPalindrome1(ch3.head.next), false)
}

func TestIsPalindrome2(t *testing.T) {
	// 1->2->3->2->1
	ch1 := newChain()
	ch1.insert(1)
	ch1.insert(2)
	ch1.insert(3)
	ch1.insert(2)
	ch1.insert(1)
	assert.Equal(t, IsPalindrome2(ch1.head.next), true)

	// 1->2->2->1
	ch2 := newChain()
	ch2.insert(1)
	ch2.insert(2)
	ch2.insert(2)
	ch2.insert(1)
	assert.Equal(t, IsPalindrome2(ch2.head.next), true)

	// 1->2->3->1->2
	ch3 := newChain()
	ch3.insert(1)
	ch3.insert(2)
	ch3.insert(3)
	ch3.insert(1)
	ch3.insert(2)
	assert.Equal(t, IsPalindrome2(ch3.head.next), false)
}

func TestIsPalindrome3(t *testing.T) {
	// 1->2->3->2->1
	ch1 := newChain()
	ch1.insert(1)
	ch1.insert(2)
	ch1.insert(3)
	ch1.insert(2)
	ch1.insert(1)
	assert.Equal(t, IsPalindrome3(ch1.head.next), true)

	// 1->2->2->1
	ch2 := newChain()
	ch2.insert(1)
	ch2.insert(2)
	ch2.insert(2)
	ch2.insert(1)
	assert.Equal(t, IsPalindrome3(ch2.head.next), true)

	// 1->2->3->1->2
	ch3 := newChain()
	ch3.insert(1)
	ch3.insert(2)
	ch3.insert(3)
	ch3.insert(1)
	ch3.insert(2)
	assert.Equal(t, IsPalindrome3(ch3.head.next), false)
}

func TestChainPartition(t *testing.T) {
	fmt.Println("===== 单链表分区 =====")
	ch1 := newChain()
	ch1.insert(9)
	ch1.insert(10)
	ch1.insert(5)
	ch1.insert(5)
	ch1.insert(6)
	ch1.insert(3)
	ch1.insert(4)
	ch1.insert(1)

	fmt.Println("===== ch1分区前 =====")
	ch1.traverse()
	fmt.Println("===== ch1分区后 =====")
	head := ChainPartition(ch1.head.next, 5)
	ch1.head.next = head
	ch1.traverse()

	ch2 := newChain()
	ch2.insert(9)
	ch2.insert(10)
	ch2.insert(5)
	ch2.insert(5)
	ch2.insert(6)

	fmt.Println("===== ch2分区前 =====")
	ch2.traverse()
	fmt.Println("===== ch2分区后 =====")
	head = ChainPartition(ch2.head.next, 5)
	ch2.head.next = head
	ch2.traverse()

	ch3 := newChain()
	ch3.insert(9)
	ch3.insert(10)
	ch3.insert(6)
	ch3.insert(3)
	ch3.insert(4)
	ch3.insert(1)

	fmt.Println("===== ch3分区前 =====")
	ch3.traverse()
	fmt.Println("===== ch3分区后 =====")
	head = ChainPartition(ch3.head.next, 5)
	ch3.head.next = head
	ch3.traverse()

	ch4 := newChain()
	ch4.insert(5)
	ch4.insert(5)
	ch4.insert(3)
	ch4.insert(4)
	ch4.insert(1)

	fmt.Println("===== ch4分区前 =====")
	ch4.traverse()
	fmt.Println("===== ch4分区后 =====")
	head = ChainPartition(ch4.head.next, 5)
	ch4.head.next = head
	ch4.traverse()

	fmt.Println("===== 单链表分区 =====")
}

func initRandChian() *randNode {
	head := newRandNode(0)
	node1 := newRandNode(1)
	node2 := newRandNode(2)
	node3 := newRandNode(3)
	node4 := newRandNode(4)

	head.next = node1

	node1.next = node2
	node1.rand = node2

	node2.next = node3
	node2.rand = node4

	node3.next = node4

	return head
}

func TestPrintRandChain(t *testing.T) {
	fmt.Println("===== 打印带有随机指针的单链表 =====")
	head := initRandChian()
	PrintRandChain(head)
	fmt.Println("===== 打印带有随机指针的单链表 =====")
}

func TestCopyListWithRand1(t *testing.T) {
	head := initRandChian()
	fmt.Println("===== 原始链表 =====")
	PrintRandChain(head)
	fmt.Println("===== 原始链表 =====")

	fmt.Println()

	headCopy := CopyListWithRand1(head)
	fmt.Println("===== 复制链表 =====")
	PrintRandChain(headCopy)
	fmt.Println("===== 复制链表 =====")
}

func TestCopyListWithRand2(t *testing.T) {
	head := initRandChian()
	fmt.Println("===== 原始链表 =====")
	PrintRandChain(head)
	fmt.Println("===== 原始链表 =====")

	fmt.Println()

	headCopy := CopyListWithRand2(head)
	fmt.Println("===== 复制链表 =====")
	PrintRandChain(headCopy)
	fmt.Println("===== 复制链表 =====")
}

func TestAddTwoNumbers(t *testing.T) {
	for i := 0; i < 5000; i++ {
		ch1, v1 := generateRandChain(9000)
		ch2, v2 := generateRandChain(9000)

		chA := generateChain(v1 + v2)

		head := AddTwoNumbers(ch1.head.next, ch2.head.next)

		chB := newChain()
		chB.head.next = head

		assert.Equal(t, compareChain(chA, chB), true)
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	for i := 0; i < 5000; i++ {
		ch1, v1 := generateRandChain(9000)
		ch2, v2 := generateRandChain(9000)

		chA := generateChain(v1 + v2)

		head := AddTwoNumbers2(ch1.head.next, ch2.head.next)

		chB := newChain()
		chB.head.next = head

		assert.Equal(t, compareChain(chA, chB), true)
	}
}

func TestReverseChain(t *testing.T) {
	ch := initChain()

	ch.traverse()

	head := ch.head.next
	ch.head.next = nil

	headNode := ReverseChain(head)

	ch.head.next = headNode

	ch.traverse()
}

func TestKSortChainMerge(t *testing.T) {
	ch1 := initChain()
	ch2 := initChain()
	ch3 := initChain()

	head := []*Node{ch1.head.next, ch2.head.next, ch3.head.next}

	root := KSortChainMerge(head)

	ch := newChain()
	ch.head.next = root

	ch.traverse()
}

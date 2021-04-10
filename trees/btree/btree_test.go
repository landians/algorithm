package btree

import (
	"fmt"
	"testing"
)

func initTree() *Node {
	head := NewNode(1)
	head.left = NewNode(2)
	head.right = NewNode(3)
	head.left.left = NewNode(4)
	head.left.right = NewNode(5)
	head.right.left = NewNode(6)
	head.right.right = NewNode(7)
	return head
}

func TestOrderRecur(t *testing.T) {
	head := initTree()
	fmt.Println("===== 前序遍历-递归 =====")
	PreOrderRecur(head)
	fmt.Println()
	fmt.Println("===== 前序遍历-递归 =====")

	fmt.Println()

	fmt.Println("===== 中序遍历-递归 =====")
	InOrderRecur(head)
	fmt.Println()
	fmt.Println("===== 中序遍历-递归 =====")

	fmt.Println()

	fmt.Println("===== 后序遍历-递归 =====")
	PosOrderRecur(head)
	fmt.Println()
	fmt.Println("===== 后序遍历-递归 =====")
}

func TestLevelOrder(t *testing.T) {
	head := initTree()
	fmt.Println("===== 层序遍历 =====")
	LevelOrder(head)
	fmt.Println()
	fmt.Println("===== 层序遍历 =====")
}

func TestPreOrderUnRecur(t *testing.T) {
	head := initTree()
	fmt.Println("===== 先序遍历-非递归 =====")
	PreOrderUnRecur(head)
	fmt.Println()
	fmt.Println("===== 先序遍历-非递归 =====")
}

func TestInOrderRecur(t *testing.T) {
	head := initTree()
	fmt.Println("===== 中序遍历-非递归 =====")
	InOrderUnRecur(head)
	fmt.Println()
	fmt.Println("===== 中序遍历-非递归 =====")
}

func TestPosOrderRecur(t *testing.T) {
	head := initTree()
	fmt.Println("===== 后序遍历-非递归 =====")
	PosOrderUnRecur(head)
	fmt.Println()
	fmt.Println("===== 后序遍历-非递归 =====")
}

func TestPrintBTree(t *testing.T) {
	head := initTree()
	fmt.Println("===== 直观打印二叉树 =====")
	PrintBTree(head, 2)
	fmt.Println()
	fmt.Println("===== 直观打印二叉树 =====")
}

func TestSearch(t *testing.T) {
	head := initTree()
	fmt.Println("===== 查询二叉树 =====")
	fmt.Println(Search(head, 2))
	fmt.Println(Search(head, 10))
	fmt.Println("===== 查询二叉树 =====")
}

func TestSerializeByPre(t *testing.T) {
	head := initTree()
	fmt.Println("===== 前序序列化二叉树-递归 =====")
	fmt.Println(SerializeByPre(head))
	fmt.Println("===== 前序序列化二叉树-递归 =====")
}

func TestSerializeByIn(t *testing.T) {
	head := initTree()
	fmt.Println("===== 中序序列化二叉树-递归 =====")
	fmt.Println(SerializeByIn(head))
	fmt.Println("===== 中序序列化二叉树-递归 =====")
}

func TestSerializeByPos(t *testing.T) {
	head := initTree()
	fmt.Println("===== 后序序列化二叉树-递归 =====")
	fmt.Println(SerializeByPos(head))
	fmt.Println("===== 后序序列化二叉树-递归 =====")
}

func TestSerializeByLevel(t *testing.T) {
	head := initTree()
	fmt.Println("===== 层序序列化二叉树-非递归 =====")
	fmt.Println(SerializeByLevel(head))
	fmt.Println("===== 层序序列化二叉树-非递归 =====")
}

func TestDeserializeByPre(t *testing.T) {
	head := initTree()
	fmt.Println("===== 前序序列化二叉树-递归 =====")
	format := SerializeByPre(head)
	fmt.Println(format)
	fmt.Println("===== 前序序列化二叉树-递归 =====")

	fmt.Println()

	fmt.Println("===== 前序-反序列化-非递归 =====")
	head = DeserializeByPre(format)
	format = SerializeByPre(head)
	fmt.Println(format)
	fmt.Println("===== 前序-反序列化-非递归 =====")
}

func TestDeserializeByLevel(t *testing.T) {
	head := initTree()
	fmt.Println("===== 层序序列化二叉树-非递归 =====")
	format := SerializeByLevel(head)
	fmt.Println(format)
	fmt.Println("===== 层序序列化二叉树-非递归 =====")

	fmt.Println()

	fmt.Println("===== 层序-反序列化-非递归 =====")
	head = DeserializeByLevel(format)
	format = SerializeByLevel(head)
	fmt.Println(format)
	fmt.Println("===== 层序-反序列化-非递归 =====")
}
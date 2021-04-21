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

func initAVl() *Node {
	head := NewNode(1)
	head.left = NewNode(2)
	head.right = NewNode(3)
	head.left.left = NewNode(4)
	head.left.right = NewNode(5)
	head.left.left.left = NewNode(6)
	return head
}

func initBST() *Node {
	head := NewNode(6)
	head.left = NewNode(4)
	head.right = NewNode(7)
	head.left.left = NewNode(3)
	head.left.right = NewNode(5)
	return head
}

func initCBT() *Node {
	head := NewNode(1)
	head.left = NewNode(2)
	head.right = NewNode(3)
	head.left.left = NewNode(4)
	head.left.right = NewNode(5)
	head.right.left = NewNode(6)
	head.right.right = NewNode(7)
	head.left.left.left = NewNode(8)
	head.left.left.right = NewNode(9)
	return head
}

func initUnCBT() *Node {
	head := NewNode(1)
	head.left = NewNode(2)
	head.right = NewNode(3)
	head.left.left = NewNode(4)
	head.left.right = NewNode(5)
	head.right.left = NewNode(6)
	head.right.right = NewNode(7)
	head.right.right.right = NewNode(8)
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

func TestIsBalanced(t *testing.T) {
	head := initTree()
	fmt.Println("===== 判断平衡二叉树 =====")
	fmt.Println(IsBalanced(head))
	fmt.Println("===== 判断平衡二叉树 =====")

	fmt.Println()

	head = initAVl()
	fmt.Println("===== 判断平衡二叉树 =====")
	fmt.Println(IsBalanced(head))
	fmt.Println("===== 判断平衡二叉树 =====")
}

func TestHeight(t *testing.T) {
	h := NewNode(100)

	fmt.Println("===== 二叉树的高度 =====")
	fmt.Println(Height(h))
	fmt.Println("===== 二叉树的高度 =====")

	fmt.Println()

	head := initTree()
	fmt.Println("===== 二叉树的高度 =====")
	fmt.Println(Height(head))
	fmt.Println("===== 二叉树的高度 =====")
}

func TestIsBST(t *testing.T) {
	head := initTree()
	fmt.Println("===== 判断搜索二叉树 =====")
	fmt.Println(IsBST(head))
	fmt.Println("===== 判断搜索二叉树 =====")

	fmt.Println()

	head = initBST()
	fmt.Println("===== 判断搜索二叉树 =====")
	fmt.Println(IsBST(head))
	fmt.Println("===== 判断搜索二叉树 =====")
}

func TestIsValidBST(t *testing.T) {
	head := initTree()
	fmt.Println("===== 判断搜索二叉树 =====")
	fmt.Println(IsValidBST(head))
	fmt.Println("===== 判断搜索二叉树 =====")

	fmt.Println()

	head = initBST()
	fmt.Println("===== 判断搜索二叉树 =====")
	fmt.Println(IsValidBST(head))
	fmt.Println("===== 判断搜索二叉树 =====")
}

func TestIsInBST(t *testing.T) {
	head := initBST()

	fmt.Println("===== 判断是否是搜索二叉树的节点 =====")
	fmt.Println(IsInBST(head, 7))
	fmt.Println("===== 判断是否是搜索二叉树的节点 =====")

	fmt.Println()

	fmt.Println("===== 判断是否是搜索二叉树的节点 =====")
	fmt.Println(IsInBST(head, 20))
	fmt.Println("===== 判断是否是搜索二叉树的节点 =====")
}

func TestInsertIntoBST(t *testing.T) {
	head := initBST()
	fmt.Println("===== 搜索二叉树插入节点 =====")
	PrintBTree(head, 2)
	fmt.Println("===== 搜索二叉树插入节点 =====")

	fmt.Println()

	fmt.Println("===== 搜索二叉树插入节点 =====")
	InsertIntoBST(head, 10)
	PrintBTree(head, 2)
	fmt.Println("===== 搜索二叉树插入节点 =====")

	fmt.Println("===== 搜索二叉树插入节点 =====")
	InsertIntoBST(head, 9)
	PrintBTree(head, 2)
	fmt.Println("===== 搜索二叉树插入节点 =====")
}

func TestIsCBT(t *testing.T) {
	head := initUnCBT()
	fmt.Println("===== 判断完全二叉树 =====")
	fmt.Println(IsCBT(head))
	fmt.Println("===== 判断完全二叉树 =====")

	fmt.Println()

	head = initCBT()
	fmt.Println("===== 判断完全二叉树 =====")
	fmt.Println(IsCBT(head))
	fmt.Println("===== 判断完全二叉树 =====")
}

func TestCountNodeNum(t *testing.T) {
	head := initCBT()
	fmt.Println("===== 统计完全二叉树的节点个数 =====")
	fmt.Println(CountNodeNum(head))
	fmt.Println("===== 统计完全二叉树的节点个数 =====")
}

func TestCountBTreeNode(t *testing.T) {
	head := initTree()
	fmt.Println("===== 统计二叉树的节点个数 =====")
	fmt.Println(CountBTreeNode(head))
	fmt.Println("===== 统计二叉树的节点个数 =====")
}

func TestFlip(t *testing.T) {
	head := initTree()
	fmt.Println("===== 翻转二叉树前 =====")
	PrintBTree(head, 2)
	fmt.Println("===== 翻转二叉树前 =====")

	fmt.Println()

	fmt.Println("===== 翻转二叉树后 =====")
	Flip(head)
	PrintBTree(head, 2)
	fmt.Println("===== 翻转二叉树后 =====")
}

func TestFlatten(t *testing.T) {
	head := initTree()
	fmt.Println("===== 拉平二叉树前 =====")
	PrintBTree(head, 2)
	fmt.Println("===== 拉平二叉树前 =====")

	fmt.Println()

	fmt.Println("===== 拉平二叉树后 =====")
	Flatten(head)
	PrintBTree(head, 2)
	fmt.Println("===== 拉平二叉树后 =====")
}

func TestCreateMaxBTree(t *testing.T) {
	arr := []int{3, 2, 1, 6, 0, 5}
	fmt.Println("===== 构造最大二叉树 =====")
	head := CreateMaxBTree(arr)
	PrintBTree(head, 2)
	fmt.Println("===== 构造最大二叉树 =====")
}
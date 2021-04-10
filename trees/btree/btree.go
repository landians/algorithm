package btree

import (
	"fmt"
	"github.com/landians/algorithm/queue"
	"github.com/landians/algorithm/stack"
	"strconv"
	"strings"
)

/*
				1
              /   \
            2      3
          /  \   /  \
         4    5 6    7

先序遍历, 指的是按照 中(当前节点) - 左 - 右 的方式进行遍历: 1 2 4 5 3 6 7
中序遍历, 指的是按照 左 - 中(当前节点) - 右 的方式进行遍历: 4 2 5 1 6 3 7
后序遍历, 指的是按照 左 - 右 - 中(当前节点) 的方式进行遍历: 4 5 2 6 7 3 1
*/

// Node 二叉树实现
type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(v int) *Node {
	return &Node{value: v}
}

const (
	defaultSize = 10
)

// PreOrderRecur 递归方式的二叉树的先序遍历, 中(当前节点) - 左 - 右
func PreOrderRecur(head *Node) {
	if head == nil {
		return
	}

	// 中(当前节点)
	fmt.Print(fmt.Sprintf("%d ", head.value))
	// 左
	PreOrderRecur(head.left)
	// 右
	PreOrderRecur(head.right)
}

// InOrderRecur 递归方式的二叉树的中序遍历, 左 - 中(当前节点) - 右
func InOrderRecur(head *Node) {
	if head == nil {
		return
	}

	// 左
	InOrderRecur(head.left)
	// 中(当前节点)
	fmt.Print(fmt.Sprintf("%d ", head.value))
	// 右
	InOrderRecur(head.right)
}

// PosOrderRecur 递归方式的二叉树的后序遍历
func PosOrderRecur(head *Node) {
	if head == nil {
		return
	}

	// 左
	PosOrderRecur(head.left)
	// 右
	PosOrderRecur(head.right)
	// 中(当前节点)
	fmt.Print(fmt.Sprintf("%d ", head.value))
}

// PreOrderUnRecur 非递归方式的二叉树的先序遍历, 需要借助栈实现
func PreOrderUnRecur(head *Node) {
	if head == nil {
		return
	}

	s := stack.New(defaultSize)

	// 根节点先入栈
	s.Push(head)

	for !s.IsEmpty() {
		// 出栈
		v := s.Pop().(*Node)

		// 打印
		fmt.Print(fmt.Sprintf("%d ", v.value))

		// 右子树节点入栈
		if v.right != nil {
			s.Push(v.right)
		}

		// 左子树节点入栈
		if v.left != nil {
			s.Push(v.left)
		}
	}
}

// InOrderUnRecur 非递归方式的二叉树的中序遍历, 需要借助栈实现
func InOrderUnRecur(head *Node) {
	if head == nil {
		return
	}

	s := stack.New(defaultSize)

	for !s.IsEmpty() || head != nil {
		if head != nil {
			// 只入栈左子树节点
			s.Push(head)

			// 访问左子树节点
			head = head.left
		} else {
			// 出栈
			head = s.Pop().(*Node)

			// 打印
			fmt.Print(fmt.Sprintf("%d ", head.value))

			// 访问右子树节点
			head = head.right
		}
	}
}

// PosOrderUnRecur 非递归方式的二叉树的后序遍历, 需要借助两个栈实现
func PosOrderUnRecur(head *Node) {
	if head == nil {
		return
	}

	// s1 用来保存遍历时的结点信息
	s1 := stack.New(defaultSize)
	// s2  用来排列后根顺序（根节点先进栈，右子树节点再进，左子树节点最后进
	s2 := stack.New(defaultSize)

	s1.Push(head)

	// 如下步骤和非递归版的前序遍历类似
	for !s1.IsEmpty() {
		// 出栈
		v := s1.Pop().(*Node)

		// 打印语句替换为入栈
		s2.Push(v)

		// 左子树节点入栈
		if v.left != nil {
			s1.Push(v.left)
		}

		// 右子树节点入栈
		if v.right != nil {
			s1.Push(v.right)
		}
	}

	// 逐个出栈, 出栈即为 左 - 右 - 中 的顺序
	for !s2.IsEmpty() {
		v := s2.Pop().(*Node)

		fmt.Print(fmt.Sprintf("%d ", v.value))
	}
}

// LevelOrder 层序遍历，需要借助队列来实现
func LevelOrder(head *Node) {
	if head == nil {
		return
	}

	q := queue.New(defaultSize)

	// 根节点先入队
	q.Enqueue(head)

	for !q.IsEmpty() {
		// 出队
		v := q.Dequeue().(*Node)

		// 打印
		fmt.Print(fmt.Sprintf("%d ", v.value))

		// 左子节点入队
		if v.left != nil {
			q.Enqueue(v.left)
		}

		// 右子节点入队
		if v.right != nil {
			q.Enqueue(v.right)
		}
	}
}

// PrintBTree 直观的打印一颗二叉树, n 表示缩进的层数，初始值为0
func PrintBTree(head *Node, n int) {
	if head == nil {
		return
	}

	// 遍历打印右子树
	PrintBTree(head.right, n+2)

	// 访问根节点
	for i := 0; i < n-1; i++ {
		fmt.Print(" ")
	}

	if n > 0 {
		fmt.Print("---")
		fmt.Println(head.value)
	}

	// 遍历打印左子树
	PrintBTree(head.left, n+2)
}

// SerializeByPre 使用递归版的先序遍历的方式来进行序列化, 使用 # 做空节点的占位符, 使用 _ 来分割节点
func SerializeByPre(head *Node) string {
	if head == nil {
		return "#_"
	}

	format := fmt.Sprintf("%d_", head.value)
	format += SerializeByPre(head.left)
	format += SerializeByPre(head.right)
	return format
}

// SerializeByIn 使用递归版的中序遍历的方式来进行序列化, 使用 # 做空节点的占位符, 使用 _ 来分割节点
func SerializeByIn(head *Node) string {
	if head == nil {
		return "#_"
	}

	format := SerializeByIn(head.left)
	format += fmt.Sprintf("%d_", head.value)
	format += SerializeByIn(head.right)
	return format
}

// SerializeByPos 使用递归版的后序遍历的方式来进行序列化, 使用 # 做空节点的占位符, 使用 _ 来分割节点
func SerializeByPos(head *Node) string {
	if head == nil {
		return "#_"
	}

	format := SerializeByPos(head.left)
	format += SerializeByPos(head.right)
	format += fmt.Sprintf("%d_", head.value)
	return format
}

// SerializeByLevel 使用层序遍历的方式来进行序列化, 使用 # 做空节点的占位符, 使用 _ 来分割节点
func SerializeByLevel(head *Node) string {
	if head == nil {
		return "#_"
	}

	format := ""

	q := queue.New(defaultSize)

	// 根节点先入队
	q.Enqueue(head)

	for !q.IsEmpty() {
		// 出队
		v := q.Dequeue().(*Node)

		// 打印
		format += fmt.Sprintf("%d_", v.value)

		// 左子节点入队
		if v.left != nil {
			q.Enqueue(v.left)
		} else {
			format += "#_"
		}

		// 右子节点入队
		if v.right != nil {
			q.Enqueue(v.right)
		} else {
			format += "#_"
		}
	}

	return format
}

// DeserializeByPre 递归版的前序序列化数据的反序列化
func DeserializeByPre(format string) *Node {
	if len(format) == 0 {
		return nil
	}

	// 得到所有的二叉树节点
	values := strings.Split(format, "_")

	// 使用一个队列来存储二叉树节点
	q := queue.New(defaultSize * 2)

	for _, v := range values {
		q.Enqueue(v)
	}
	return recoverByQueue(q)
}

// 通过队列来还原树
func recoverByQueue(q *queue.Queue) *Node {
	if q == nil || q.IsEmpty() {
		return nil
	}

	// 出队
	v := q.Dequeue().(string)
	if v == "#" {
		return nil
	}

	n, _ := strconv.Atoi(v)
	head := NewNode(n)
	head.left = recoverByQueue(q)
	head.right = recoverByQueue(q)
	return head
}

// DeserializeByLevel 层序序列化数据的反序列化
func DeserializeByLevel(format string) *Node {
	if len(format) == 0 {
		return nil
	}

	// 得到所有的二叉树节点
	values := strings.Split(format, "_")
	i := 0

	// 得到根节点，并入队
	head := createNode(values[i])
	i++
	q := queue.New(defaultSize * 2)
	q.Enqueue(head)

	var node *Node
	for !q.IsEmpty() {
		// 出队当前子树的根节点
		node = q.Dequeue().(*Node)

		// 赋值左子树节点
		node.left = createNode(values[i])
		i++

		// 赋值右子树节点
		node.right = createNode(values[i])
		i++

		// 左子树节点入队
		if node.left != nil {
			q.Enqueue(node.left)
		}

		// 右子树节点入队
		if node.right != nil {
			q.Enqueue(node.right)
		}
	}
	return head
}

func createNode(v string) *Node {
	if v == "#" {
		return nil
	}
	n, _ := strconv.Atoi(v)
	return NewNode(n)
}

// Search 查询二叉树中的节点
func Search(head *Node, v int) *Node {
	if head == nil {
		return nil
	}

	var find *Node

	// 按照 中(当前节点) - 左 - 右 的顺序查询
	if head.value == v {
		find = head
	} else {
		if find = Search(head.left, v); find == nil {
			find = Search(head.right, v)
		}
	}

	return find
}

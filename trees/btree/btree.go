package btree

import (
	"fmt"
	"github.com/landians/algorithm/queue"
	"github.com/landians/algorithm/stack"
	"math"
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

// Node 二叉树节点的实现
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
		// 出栈, 顺序为：中 - 右 - 左
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
	return recoverByPreQueue(q)
}

// 通过队列来还原树
func recoverByPreQueue(q *queue.Queue) *Node {
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
	head.left = recoverByPreQueue(q)
	head.right = recoverByPreQueue(q)
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

// MaxWidth 返回二叉树的最大宽度
func MaxWidth(head *Node) int {
	if head == nil {
		return 0
	}

	q := queue.New(defaultSize)
	levelMap := make(map[*Node]int) // 用于记录每个节点所在的层数
	levelMap[head] = 1              // 根节点是在第一层

	curLevel := 1        // 记录当前的层数
	curLevelNodes := 0   // 记录当前层的节点个数
	max := math.MinInt64 // 记录所有层中，哪一层的节点个数是最多的，其实也就是二叉树的最大宽度

	// 根节点先入队
	q.Enqueue(head)

	for !q.IsEmpty() {
		// 出队
		v := q.Dequeue().(*Node)

		// 得到节点所在的层数
		curNodeLevel := levelMap[v]
		if curNodeLevel == curLevel { // 弹出的节点所在的层正好是统计层
			curLevelNodes++
		} else { // 说明 curLevel 层已经统计完了，curLevelNodes 的值就是 curLevel 的宽度
			// 进入下一层继续统计
			curLevel++
			// 进入新的一层的时候，curLevelNodes 自然为1
			// 因为出队的节点必然是新的一层的第一个节点
			curLevelNodes = 1
		}

		// 左子节点入队
		if v.left != nil {
			levelMap[v.left] = curNodeLevel + 1
			q.Enqueue(v.left)
		}

		// 右子节点入队
		if v.right != nil {
			levelMap[v.right] = curNodeLevel + 1
			q.Enqueue(v.right)
		}

		// 更新二叉树的最大宽度值
		if max < curLevelNodes {
			max = curLevelNodes
		}
	}

	return max
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

/*
	平衡二叉树的定义: 对于二叉树中的任何一个节点, 其左子树节点的高度 和 右子树节点的高度差不超过1.

	如果判断一颗二叉树是否为平衡二叉树？
	将问题分解为：只要保证以二叉树的每个结点为根节点的树是否平衡；而遍历到每个结点时，要想知道以该结点为根结点的子树是否是平衡二叉树，
	首先判断该结点的左子树是否平衡，然后判断该结点的右子树是否平衡，如果两个子树都平衡，分别计算左子树和右子树的高度，因此主要需要
	以下两个信息:
	(1) 该结点的左子树、右子树是否是平衡二叉树
	(2) 左右子树的高度分别是多少，相差是否超过 1
	可以使用递归实现

	eg: 平衡二叉树
			    1
              /   \
            2      3
          /  \   /  \
         4    5 6    7

		非平衡二叉树: 左子树高度 = 3, 右子树高度 = 1; 3 - 1 = 2 > 1
				1
              /  \
            2     3
          /  \
         4    5
		/
       6
*/

// IsBalanced 用于判断二叉树是否平衡，返回二叉树是否平衡和二叉树的高度
func IsBalanced(head *Node) bool {
	isB, _ := processBTree(head)
	return isB
}

// Height 获取二叉树的高度
func Height(head *Node) int {
	if head == nil {
		return 0
	}

	// 获取左子树的高度
	lh := Height(head.left)

	// 获取右子树的高度
	rh := Height(head.left)

	// 取左右子树的高度的最大值 +  1
	return int(math.Max(float64(lh), float64(rh))) + 1
}

// 返回二叉树是否平衡和二叉树的高度
func processBTree(head *Node) (bool, int) {
	// 空树是平衡的
	if head == nil {
		return true, 0
	}

	// 判断左子树是否平衡，若不平衡，则返回的高度默认取0
	isB, lh := processBTree(head.left)
	if !isB {
		return false, 0
	}

	// 判断右子树是否平衡，若不平衡，则返回的高度默认取0
	isB, rh := processBTree(head.right)
	if !isB {
		return false, 0
	}

	// 若左子树和右子树都平衡，则判断高度差是否大于1
	if math.Abs(float64(lh-rh)) > 1 {
		return false, 0
	}

	// 返回给上一层的高度(最底层是平衡的), 取左右子树的高度的最大值 +  1
	return true, int(math.Max(float64(lh), float64(rh))) + 1
}

/*
	搜索二叉树的定义: 对于二叉树的任意一个节点，其左子树节点的值总是小于其根节点的值，其右子树节点的值总是大于其根节点的值
	eg:
			6
          /   \
         4    7
       /  \
      3    5
    如果判断是否为搜索二叉树?
    搜索二叉树有一个特点: 其中序遍历得到的节点的值的集合是升序排列的, 因此可以使用非递归版的中序遍历来进行判断

	搜索二叉树的重复值怎么处理？
	可以定义节点的结构为:
	type Node struct {
		Value int
		n int
	}
	其中 n 为值等于 value 的节点的数量，相当于吧重复的节点给压缩了。
*/
func IsBST(head *Node) bool {
	if head == nil {
		return false
	}

	s := stack.New(defaultSize)

	// 表示中序遍历过程中的上一个节点的值
	compare := math.MinInt64

	for !s.IsEmpty() || head != nil {
		if head != nil {
			// 只入栈左子树节点
			s.Push(head)

			// 访问左子树节点
			head = head.left
		} else {
			// 出栈
			head = s.Pop().(*Node)

			// 进行比较
			if head.value < compare {
				return false
			} else {
				compare = head.value
			}

			// 访问右子树节点
			head = head.right
		}
	}
	return true
}

// IsValidBST 使用递归的方式来判断是否为 BST
func IsValidBST(head *Node) bool {
	return isValidBST(head, nil, nil)
}

// 任意一个节点，都必须满足，max.Value > head.Value > min.Value
func isValidBST(head *Node, min *Node, max *Node) bool {
	if head == nil {
		return true
	}

	if min != nil && head.value < min.value {
		return false
	}

	if max != nil && head.value > max.value {
		return false
	}

	return isValidBST(head.left, min, head) && isValidBST(head.right, head, max)
}

// IsInBST 在 BST 内搜索指定节点
func IsInBST(head *Node, target int) bool {
	if head == nil {
		return false
	}

	if head.value == target {
		return true
	}

	if head.value > target {
		return IsInBST(head.left, target)
	} else {
		return IsInBST(head.right, target)
	}
}

/*
	抽象 BST 的遍历框架
	func BST(head *Node, target interface) interface{} {
		if head == nil {
			return nil
		}

		if head.value == target {
			... do something
		}

		if head.value > target {
			return BST(head.left, target)
		} else {
			return BST(head.right, target)
		}
	}
*/

// InsertIntoBST 插入节点到 BST 中
func InsertIntoBST(head *Node, v int) *Node {
	// 找到空位置插入新的节点
	if head == nil {
		return NewNode(v)
	}

	// if (head.value == v)
	// BST 中一般不会插入已存在元素

	// 左子树查找
	if head.value < v {
		head.right = InsertIntoBST(head.right, v)
	}

	// 右子树查找
	if head.value > v {
		head.left = InsertIntoBST(head.left, v)
	}

	return head
}

// DeleteFromBST 从 BST 中删除指定节点
func DeleteFromBST(head *Node, target int) *Node {
	if head == nil {
		return nil
	}

	if head.value == target {
		// 执行删除操作
	}

	// 左子树查找
	if head.value > target {
		head.left = DeleteFromBST(head.left, target)
	}

	// 右子树查找
	if head.value < target {
		head.right = DeleteFromBST(head.right, target)
	}

	return head
}

/*
	完全二叉树举例:
						1
                     /     \
                   2       5
                 /   \   /   \
                3     4 6    7
               /
              8

	如何判断是否为完全二叉树？ 需要使用层序遍历来判断
	1. 如果一个二叉树节点，有右子树，但没有左子树，一定不是完全二叉树。
		1
		 \
          x
	2. 如果一个二叉树节点有左子树但没有右子树或者左右子树都没有，则开启一个开关，接下来每个遍历的二叉树节点都必须为叶子节点
		1			1
      /  \   或    /  \
     x    x       2   x

*/
func IsCBT(head *Node) bool {
	if head == nil {
		return true
	}

	// 用于开启是否需要判断二叉树节点为叶子节点的开关
	leaf := false

	q := queue.New(defaultSize * 2)

	// 根节点先入队
	q.Enqueue(head)

	for !q.IsEmpty() {
		// 出队
		v := q.Dequeue().(*Node)

		// leaf = true 时, 需要判断每个二叉树节点是否为叶子节点，如果不是，则不是完全二叉树
		if leaf && !isLeaf(v) {
			return false
		}

		// 有右子树但没有左子树，则不是完全二叉树
		if v.right != nil && v.left == nil {
			return false
		}

		// 二叉树节点有左子树但没有右子树或者左右子树都没有，则开启是否为叶子节点的开关判断
		if (v.left != nil && v.right == nil) || (v.left == nil && v.right == nil) {
			leaf = true
		}

		// 左子节点入队
		if v.left != nil {
			q.Enqueue(v.left)
		}

		// 右子节点入队
		if v.right != nil {
			q.Enqueue(v.right)
		}
	}

	return true
}

// 判断二叉树节点是否为叶子节点
func isLeaf(head *Node) bool {
	if head == nil {
		return true
	}

	return head.left == nil && head.right == nil
}

/*
	获取完全二叉树的节点的个数，要求时间复杂度为O(N), 因此不能使用遍历的方式来计算。
	大致计算流程分为两个步骤:
	1. 判断左子树/右子树为满二叉树, 从而利用满二叉树的节点的个数 = 2 ^ h - 1, 其中 h 为树高, 这个公式来计算得到左子树/右子树的节点个数
	2. 遍历另一半的二叉树来计算得到节点个数。

	详细计算流程如下:
	1. 遍历到左子树的叶子节点，得到左子树的树高h1;
	2. 遍历到根节点的右节点的左子树的叶子节点, 同样得到右子树的树高h2;
	3. 若 h2 < h1, 则说明右子树是一颗满二叉树, 否则 h1 = h2 左子树是一颗满二叉树;
    4. 根据满二叉树的计算节点的公式得到左子树/右子树的节点的个数 n1;
	5. 遍历另一半的二叉树, 并计算得到节点的个数 n2;
    6. 计算总的节点的个数: n =  n1 + n2 + 1
	eg:
					1
				/       \
               2         3
            /    \     /   \		h1 = 3, h2 = 2; h2 < h1,  右子树为满二叉树
           4      5   6     7
         /   \
         8   9

					1
				/       \
               2          3
            /    \      /   \		h1 = 3, h2 = 2; h2 > h1,  左子树为满二叉树
           4      5    6     7
         /   \   / \  /
         8   9  10 11 12
*/
func CountNodeNum(head *Node) int {
	if head == nil {
		return 0
	}
	return doCountNodeNum(head, 1, mostLeftLevel(head, 1))
}

// 返回节点个数, node 为当前节点,  level 为当前节点所在层数, h 为这个完全二叉树层数
func doCountNodeNum(node *Node, level int, h int) int {
	if level == h {
		return 1
	}

	// h1 = h2， 完全二叉树中的每一个子树都是完全二叉树
	if mostLeftLevel(node.right, level+1) == h {
		return (1 << (h - level)) + doCountNodeNum(node.right, level+1, h)
	} else {
		return (1 << (h - level - 1)) + doCountNodeNum(node.left, level+1, h)
	}
}

// 遍历当前节点的左子树至叶子节点, 计算高度
func mostLeftLevel(node *Node, level int) int {
	for node != nil {
		level++
		node = node.left
	}
	return level - 1
}

// CountBTreeNode 计算二叉树的节点的个数
func CountBTreeNode(head *Node) int {
	if head == nil {
		return 0
	}
	// 总节点的个数 = 根节点 + 左子树节点个数 + 右子树的节点个数
	return 1 + CountBTreeNode(head.left) + CountBTreeNode(head.right)
}

/*
	Flip 翻转二叉树:
	翻转前:
			    1
              /   \
            2      3
          /  \   /  \
         4    5 6    7
	翻转后:
			    1
              /   \
            3      2
          /  \   /  \
         7   6  5    4

*/
func Flip(head *Node) {
	if head == nil {
		return
	}

	// 获取左子树
	left := head.left
	// 获取右子树
	right := head.right

	// 交换指向
	head.left = right
	head.right = left

	// 对于每一个二叉树的节点都这么做
	Flip(head.left)
	Flip(head.right)
}

/*
	Flatten 二叉树按照 中 - 左 - 右的方式展开为链表
		1                 1              1
      /  \			    /  \              \
     2    3		=>     2   3     =>        2
   /  \  / \            \   \               \
  4   5 6  7             4   6               4
                          \   \               \
                          5    7              5
                                                \
                                                 3
                                                  \
                                                   6
      											    \
                                                     7
*/
func Flatten(head *Node) {
	if head == nil {
		return
	}

	// 先拉平左右子树
	Flatten(head.left)
	Flatten(head.right)

	// 左右子树已经被拉平成一条链表
	left := head.left
	right := head.right

	// 将左子树作为右子树
	head.left = nil
	head.right = left

	// 将原先的右子树接到当前右子树的末端
	cur := head
	for cur.right != nil {
		cur = cur.right
	}
	cur.right = right
}

/*
	CreateMaxBTree: 根据数组元素，构建最大的二叉树，即对于二叉树的任意一个节点，其值都大于它左子树的节点的值和右子树的节点的值
	eg: [3, 2, 1, 6, 0, 5]
				6
			  /  \
			 3    5
             \   /
             2  0
              \
              1
*/

func CreateMaxBTree(arr []int) *Node {
	return build(arr, 0, len(arr)-1)
}

func max(arr []int, lo, hi int) (int, int) {
	maxValue := math.MinInt64
	index := -1
	for i := lo; i <= hi; i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
			index = i
		}
	}
	return maxValue, index
}

func build(arr []int, lo, hi int) *Node {
	if len(arr) == 0 {
		return nil
	}

	if lo > hi {
		return nil
	}

	// 先找到数组中的最大值和其对应的索引
	maxValue, i := max(arr, lo, hi)

	fmt.Printf("[DEBUG] max = %d, index = %d\n", maxValue, i)

	// 根节点
	head := NewNode(maxValue)

	// 构造最大左子树
	head.left = build(arr, lo, i-1)

	// 构造最大右子树
	head.right = build(arr, i+1, hi)

	return head
}

// LowestCommonAncestor 返回两个节点的最低公共祖先节点,  node1 和 node2 肯定要是以 head 为根节点的二叉树上的节点
func LowestCommonAncestor(head *Node, node1 *Node, node2 *Node) *Node {
	if head == nil || node1 == nil || node2 == nil {
		return nil
	}

	// 用于记录每一个节点的父节点
	fatherMap := make(map[*Node]*Node)
	// 先搞定根节点
	fatherMap[head] = head
	// 再搞定左右子树
	processLca(head, fatherMap)

	// 用于记录 node1 网上寻找父节点的路径
	node1Map := make(map[*Node]struct{})
	cur := node1

	// node1 开始向上寻找自己的父节点，并记录寻找路径
	for cur != fatherMap[cur] {
		node1Map[cur] = struct{}{}
		cur = fatherMap[cur]
	}
	node1Map[head] = struct{}{}

	// node2 开始向上寻找自己的父节点
	cur = node2
	for cur != fatherMap[cur] {
		// 当发现 node2 的其中一个父节点也是 node1 的父节点时，就说明找到了 node1 和 node2 的最低公共祖先节点了
		if _, ok := node1Map[cur]; ok {
			return cur
		}
		cur = fatherMap[cur]
	}
	return cur
}

func processLca(head *Node, fatherMap map[*Node]*Node) {
	if head == nil {
		return
	}

	if head.left != nil {
		fatherMap[head.left] = head
		processLca(head.left, fatherMap)
	}

	if head.right != nil {
		fatherMap[head.right] = head
		processLca(head.right, fatherMap)
	}
}

func LowestCommonAncestorRecur(head *Node, node1 *Node, node2 *Node) *Node {
	if head == nil || head == node1 || head == node2 {
		return head
	}

	// 获取左子树上两个节点的最低公共祖先节点
	left := LowestCommonAncestorRecur(head.left, node1, node2)
	// 获取右子树上两个节点的最低公共祖先节点
	right := LowestCommonAncestorRecur(head.right, node1, node2)

	// 说明 node1 和 node2 不是对方的 lca
	if left != nil && right != nil {
		return head
	}

	// 说明 node1 或者 node2 是对方的 lca
	if left != nil {
		return left
	}
	return right
}

// Morris Morris 遍历一个二叉树
func Morris(head *Node) {
	if head == nil {
		return
	}

	var (
		cur       *Node = head
		mostRight *Node
	)

	for cur != nil {
		// 这里根据打印结果可知, 如果当前节点是存在左子树的, 则会被打印两遍
		fmt.Print(fmt.Sprintf("%d ", cur.value))
		// 判断当前节点是否存在左子树
		mostRight = cur.left
		if mostRight != nil {
			// 若当前节点存在左子树, 则去查找左子树的真实的右节点, 这里的真实指的是 mostRight != cur
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}

			if mostRight.right == nil { // 若查找到的真实的右节点的右子树为空, 则其右子树指向当前节点, 同时当前节点指向其左左子树节点
				mostRight.right = cur
				cur = cur.left
				continue // continue 很重要
			} else { // 若查找到的真实的右节点的右子树不为空, 这里一般是 mostRight.right == cur 了, 为了不陷入死循环, 这里 mostRight.right = nil
				mostRight.right = nil
			}
		}
		// 当前节点不存在左子树时, 当前节点指向其右子树, 这一步也是 mostRight.right = nil 后会执行的关键一步, 防止陷入了死循环
		cur = cur.right
	}
}

/*
MorrisPreOrder
Morris 实现先序遍历一个二叉树
打印原则说明: 存在左子树的节点, 第一次遍历到时打印, 第二次遍历到时不打印, 不存在左子树的节点直接打印
*/
func MorrisPreOrder(head *Node) {
	if head == nil {
		return
	}

	var (
		cur       *Node = head
		mostRight *Node
	)

	for cur != nil {
		// 判断当前节点是否存在左子树
		mostRight = cur.left
		if mostRight != nil {
			// 若当前节点存在左子树, 则去查找左子树的真实的右节点, 这里的真实指的是 mostRight != cur
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}

			if mostRight.right == nil { // 若查找到的真实的右节点的右子树为空, 则其右子树指向当前节点, 同时当前节点指向其左左子树节点
				mostRight.right = cur
				// 存在左子树的节点, 第一次遍历到时打印
				fmt.Print(fmt.Sprintf("%d ", cur.value))
				cur = cur.left
				continue // continue 很重要
			} else { // 若查找到的真实的右节点的右子树不为空, 这里一般是 mostRight.right == cur 了, 为了不陷入死循环, 这里 mostRight.right = nil
				mostRight.right = nil
			}
		} else {
			// 不存在左子树的节点直接打印
			fmt.Print(fmt.Sprintf("%d ", cur.value))
		}

		// 当前节点不存在左子树时, 当前节点指向其右子树, 这一步也是 mostRight.right = nil 后会执行的关键一步, 防止陷入了死循环
		cur = cur.right
	}
}

/*
MorrisInOrder
Morris 实现中序遍历一个二叉树
打印原则说明: 存在左子树的节点, 第一次遍历到时不打印, 第二次遍历到时打印, 不存在左子树的节点直接打印
*/
func MorrisInOrder(head *Node) {
	if head == nil {
		return
	}

	var (
		cur       *Node = head
		mostRight *Node
	)

	for cur != nil {
		// 判断当前节点是否存在左子树
		mostRight = cur.left
		if mostRight != nil {
			// 若当前节点存在左子树, 则去查找左子树的真实的右节点, 这里的真实指的是 mostRight != cur
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}

			if mostRight.right == nil { // 若查找到的真实的右节点的右子树为空, 则其右子树指向当前节点, 同时当前节点指向其左左子树节点
				mostRight.right = cur
				cur = cur.left
				continue // continue 很重要
			} else { // 若查找到的真实的右节点的右子树不为空, 这里一般是 mostRight.right == cur 了, 为了不陷入死循环, 这里 mostRight.right = nil
				mostRight.right = nil
			}
		}

		fmt.Print(fmt.Sprintf("%d ", cur.value))
		// 当前节点不存在左子树时, 当前节点指向其右子树, 这一步也是 mostRight.right = nil 后会执行的关键一步, 防止陷入了死循环
		cur = cur.right
	}
}

package _34_copyRandomList

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	p := head

	// 使用一个 map 存储 节点与克隆节点之间的关系
	copys := make(map[*Node]*Node)
	for p != nil {
		cp := &Node{Val: p.Val}
		copys[p] = cp
		p = p.Next
	}

	// 补齐 next 指针和 rand 指针的指向
	p = head
	for p != nil {
		copys[p].Next = copys[p.Next]
		copys[p].Random = copys[p.Random]
		p = p.Next
	}

	return copys[head]
}
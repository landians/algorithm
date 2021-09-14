package graph

import (
	"container/heap"
	"container/list"
	"fmt"
	"github.com/landians/algorithm/queue"
	"github.com/landians/algorithm/stack"
)

// Graph 图的抽象
type Graph struct {
	Nodes map[int]*Node      // 点集合
	Edges map[*Edge]struct{} // 边集合
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
		Edges: make(map[*Edge]struct{}),
	}
}

// Node 点的抽象
type Node struct {
	Value int     // 数据项
	In    int     // 被其他点指向的边的数量，也称为这个点的出度
	Out   int     // 指向其他点的边的数量，也称为这个点的入度
	Nexts []*Node // 指向的点的集合
	Edges []*Edge // 指向的边的集合
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		In:    0,
		Out:   0,
		Nexts: make([]*Node, 0),
		Edges: make([]*Edge, 0),
	}
}

// Edge 边的抽象
type Edge struct {
	Weight int   // 权重值
	From   *Node // 起始节点
	To     *Node // 结束节点
}

func NewEdge(weight int, from *Node, to *Node) *Edge {
	return &Edge{
		Weight: weight,
		From:   from,
		To:     to,
	}
}

/*
CreateGraph 根据一系列数组结构来构建图
数组结构说明：[权重, from节点, to节点], 例如: [3, 2, 4], 表示节点 2 的邻接节点是 节点 4, 边的权重为 3
*/
func CreateGraph(matrix [][]int) *Graph {
	g := NewGraph()

	for i := 0; i < len(matrix); i++ {
		weight := matrix[i][0]
		from := matrix[i][1]
		to := matrix[i][2]

		// 若在 graph 中没有 from 这个点则建上这个点
		if _, ok := g.Nodes[from]; !ok {
			g.Nodes[from] = NewNode(from)
		}

		// 若在 graph 中没有 to 这个点则建上这个点
		if _, ok := g.Nodes[to]; !ok {
			g.Nodes[to] = NewNode(to)
		}

		// 取出 from 节点和 to 节点，建立新的边
		fromNode := g.Nodes[from]
		toNode := g.Nodes[to]
		edge := NewEdge(weight, fromNode, toNode)

		// 将 to 节点加入到 from 节点的邻接节点集合
		fromNode.Nexts = append(fromNode.Nexts, toNode)
		// from 节点的出度+1
		fromNode.Out++
		// to 节点的入度+1
		toNode.In++
		// 将 edge 加入到 from 节点的边集合中
		fromNode.Edges = append(fromNode.Edges, edge)
		// 将 edge 加入到整个图的边的集合中
		g.Edges[edge] = struct{}{}
	}
	return g
}

// BFS 图的广度遍历搜索
func BFS(node *Node) {
	if node == nil {
		return
	}

	q := queue.New(10)
	filter := make(map[*Node]struct{})

	q.Enqueue(node)
	filter[node] = struct{}{}

	for !q.IsEmpty() {
		curNode := q.Dequeue().(*Node)
		// 此处就是处理遍历到的节点的代码
		fmt.Println("node value:", curNode.Value)
		for _, next := range curNode.Nexts {
			if _, ok := filter[next]; !ok {
				q.Enqueue(next)
				filter[next] = struct{}{}
			}
		}
	}
}

// DFS 图的深度优先遍历搜索
func DFS(node *Node) {
	if node == nil {
		return
	}

	s := stack.New(10)
	filter := make(map[*Node]struct{})

	s.Push(node)
	filter[node] = struct{}{}

	for !s.IsEmpty() {
		curNode := s.Pop().(*Node)
		fmt.Println("node value:", curNode.Value)
		for _, next := range curNode.Nexts {
			if _, ok := filter[next]; !ok {
				s.Push(curNode)
				s.Push(next)
				filter[next] = struct{}{}
				fmt.Println("node value:", next.Value)
				break
			}
		}
	}
}

func SortTopology(g *Graph) {
	if g == nil {
		return
	}

	in := make(map[*Node]int) // key 为节点，value 表示该节点剩余的入度值
	q := queue.New(10)

	// 将第一个入度为0的节点入队，且只有入度为0的节点才能入队
	for _, node := range g.Nodes {
		// 记录所有图的节点的真实入度值
		in[node] = node.In
		if node.In == 0 {
			q.Enqueue(node)
		}
	}

	//  拓扑排序的结果，依次加入到 arr
	arr := make([]*Node, 0, 10)
	for !q.IsEmpty() {
		curNode := q.Dequeue().(*Node)
		arr = append(arr, curNode)

		// 抹去节点的影响
		for _, next := range curNode.Nexts {
			in[next] -= 1
			if in[next] == 0 {
				q.Enqueue(next)
			}
		}
	}

	// 这里用于处理排序后的数组
	for _, v := range arr {
		fmt.Println(v.Value)
	}
}

// 非并查集实现判断图是否有环

type NodeSet []*Node

// MySets 中的 SetMap 用于记录每个图的节点所对应的集合的映射关系
type MySets struct {
	SetMap map[*Node]NodeSet
}

func NewMySet(nodes []*Node) *MySets {
	mySets := &MySets{
		SetMap: make(map[*Node]NodeSet),
	}
	for _, node := range nodes {
		set := NodeSet(make([]*Node, 0))
		set = append(set, node)
		mySets.SetMap[node] = set
	}
	return mySets
}

// IsSameSet 判断两个图的节点所在的集合是否为同一个集合, 通过判断地址是否相同即可
func (s *MySets) IsSameSet(from *Node, to *Node) bool {
	fromSet := s.SetMap[from]
	toSet := s.SetMap[to]
	return &fromSet == &toSet
}

// Union 用于将两个图的节点所在的集合合并
func (s *MySets) Union(from *Node, to *Node) {
	fromSet := s.SetMap[from]
	toSet := s.SetMap[to]

	for _, toNode := range toSet {
		// 将 toSet 中的所有 Node 加入到 fromSet 中
		fromSet = append(fromSet, toNode)
		// 将 toNode 所在的集合更改 fromSet
		s.SetMap[toNode] = fromSet
	}
}


func PrimMST(g *Graph) {
	// TODO: 创建优先级队列，本质是一个小顶堆
	priorityQueue := heap.New()

	// 新建用于保存所有 Node 的 map
	nodeMap := make(map[*Node]struct{})

	// 新建用于保存最终确定的最小生成树的边的 map
	mstMap := make(map[*Edge]struct{})

	// for 循环是为了考虑到 森林 的情况，即多个图一起构建的一个数据结构
	for _, node := range g.Nodes { // 随机选择一个 node

		// 节点不存在，则新添加节点
		if _, ok := nodeMap[node]; !ok {
			nodeMap[node] = struct{}{}
			// 并将这个节点所关联的边加入到优先级队列中
			for _, edge := range node.Edges {
				// TODO: 将边加入到优先级队列中
			}

			// 但优先级队列不为空时
			for !heap.IsEmpay() {
				// TODO: 出队一个边的权重值最小的
				edge := heap.Pop().(*Edge)
				// 选择一个节点
				toNode := edge.To
				// 若这个节点不在 nodeMap 中，则新增
				if _, ok = nodeMap[toNode]; !ok {
					nodeMap[toNode] = struct{}{}
					// 同时将这条边记录到 mstMap 中
					mstMap[edge] = struct{}{}
					// 继续加入其他的边
					for _, nextEdge := toNode.Edges {
						// TODO: 将边加入到优先级队列中
					}
				}
			}
		}
	}
}
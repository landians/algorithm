package graph

//  AlGraph 邻接表的抽象
type AlGraph struct {
	TopNodes []*TopNode // 顶点的集合
}

// ArcNode 邻接节点
type ArcNode struct {
	Index  int      // 邻接节点在顶点数组中的位置下标
	Weight int      // 邻接节点与顶点相连的边的权重值
	Next   *ArcNode // 指向的下一个邻接节点
}

// TopNode 顶点
type TopNode struct {
	Value int      // 数据项
	First *ArcNode // 指向第一个邻接节点的指针
}

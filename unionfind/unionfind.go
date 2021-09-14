package unionfind

import "github.com/landians/algorithm/stack"

type element struct {
	value int
}

// 并查集结构
type UnionFindSet struct {
	// key: 某个元素的值 value: 元素
	ElementMap map[int]*element
	// key: 某个元素 value: 该元素对应的父元素
	FatherMap map[*element]*element
	// key: 某个元素的代表元素 value: 代表元素所在集合的大小
	SizeMap map[*element]int
}

func New(values []int) *UnionFindSet {
	set := &UnionFindSet{
		ElementMap: make(map[int]*element),
		FatherMap:  make(map[*element]*element),
		SizeMap:    make(map[*element]int),
	}

	// 初始化时每个元素都是自己的集合
	for _, v := range values {
		ele := &element{value: v}
		set.ElementMap[v] = ele
		set.FatherMap[ele] = ele
		set.SizeMap[ele] = 1
	}

	return set
}

// 找到 ele 元素所在集合的代表元素, 时间复杂度接近于 O(1)
func (u *UnionFindSet) findHead(ele *element) *element {
	s := stack.New(100)

	// 向上查找父元素的过程
	for fEle := u.FatherMap[ele]; fEle != ele; ele = fEle {
		s.Push(fEle)
	}

	// 优化，使得下一次查找代表元素能直接查询到
	for !s.IsEmpty() {
		e := s.Pop().(*element)
		u.FatherMap[e] = ele
	}

	return ele
}

// IsSameSet 判断两个值所在的集合是否为同一个集合, 时间复杂度接近于 O(1)
func (u *UnionFindSet) IsSameSet(v1 int, v2 int) bool {
	ele1, ok1 := u.ElementMap[v1]
	ele2, ok2 := u.ElementMap[v2]

	if ok1 && ok2 {
		return u.findHead(ele1) == u.findHead(ele2)
	}

	return false
}

// Union 将两个值所在的集合合并成同一个集合, 时间复杂度接近于 O(1)
func (u *UnionFindSet) Union(v1 int, v2 int) {
	ele1, ok1 := u.ElementMap[v1]
	ele2, ok2 := u.ElementMap[v2]

	if ok1 && ok2 {
		// 分别找到两个元素所在集合的代表元素
		v1F := u.findHead(ele1)
		v2F := u.findHead(ele2)

		// 比较两个代表元素所在集合的元素个数
		if v1F != v2F {
			v1FSize := u.SizeMap[v1F]
			v2FSize := u.SizeMap[v2F]

			var (
				big *element
				small *element
			)

			if v1FSize >= v2FSize {
				big = v1F
				small = v2F
			} else {
				big = v2F
				small = v1F
			}

			// 将较小的集合指向较大的集合
			u.FatherMap[small] = big
			u.SizeMap[big] = v1FSize + v2FSize
			delete(u.SizeMap, small)
		}
	}
}
package heap

import "fmt"

// Heap 本身其实就是一个 slice, 也是数组形式的完全二叉树
type Heap struct {
	data     []int
	capacity int
	size     int  // 堆中元素的个数
	isMax    bool // 是否为一个大顶堆
}

func NewHeap(capacity int, isMax bool) *Heap {
	return &Heap{
		data:     make([]int, capacity),
		capacity: capacity,
		size:     0,
		isMax:    isMax,
	}
}

func NewHeapFrom(data []int, isMax bool) *Heap {
	heap := &Heap{
		data:     make([]int, len(data)),
		capacity: len(data),
		isMax:    isMax,
	}

	for _, v := range data {
		heap.Push(v)
	}

	return heap
}

func (h *Heap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) LeftChild(i int) int {
	return 2*i + 1
}

func (h *Heap) RightChild(i int) int {
	return 2*i + 2
}

func (h *Heap) Push(value int) {
	if h.size == h.capacity {
		return
	}

	h.data[h.size] = value
	h.size++
	h.shiftUp(h.size - 1)
}

func (h *Heap) Pop() int {
	if h.size == 0 {
		return -1
	}

	top := h.data[0]
	h.data[0] = h.data[h.size-1]
	h.size--
	h.shiftDown(0)
	return top
}

func (h *Heap) Data() []int {
	return h.data
}

func (h *Heap) Format() string {
	if h.size == 0 {
		return ""
	}

	heapFormat := "heap"
	for _, v := range h.data {
		heapFormat += fmt.Sprintf("->%d", v)
	}
	return heapFormat
}

// 向上进行调整
func (h *Heap) shiftUp(i int) {
	for i > 0 && h.isShiftUp(i) {
		h.data[i], h.data[h.Parent(i)] = h.data[h.Parent(i)], h.data[i]
		i = h.Parent(i)
	}
}

func (h *Heap) isShiftUp(i int) bool {
	// 对于大顶堆来说，如果子节点大于父节点的值，则需要进行位置交换
	maxShift := h.isMax && h.data[i] > h.data[h.Parent(i)]
	// 对于小顶堆来说，如果子节点小于父节点的值，则需要进行位置交换
	minShift := !h.isMax && h.data[i] < h.data[h.Parent(i)]
	return maxShift || minShift
}

// 向下进行调整
func (h *Heap) shiftDown(i int) {
	for h.LeftChild(i) < h.size {
		// 父节点需要依次和左右子节点进行比较，然后选择符合交换条件的节点进行交换， 初始化的时候先假设左子节点是符合条件的
		shiftI := h.LeftChild(i)
		rightI := h.RightChild(i)

		// 判断右子节点是否满足条件，如果满足，则将右子节点作为接下来需要交换的节点
		if rightI < h.size && h.isShiftDown(rightI, shiftI) {
			shiftI = rightI
		}

		// 不需要往下继续调整了，则 break
		if h.isShiftBreak(i, shiftI) {
			break
		}

		// 执行交换操作
		h.data[i], h.data[shiftI] = h.data[shiftI], h.data[i]
		i = shiftI
	}
}

func (h *Heap) isShiftDown(i, j int) bool {
	maxShift := h.isMax && h.data[i] > h.data[j]
	minShift := !h.isMax && h.data[i] < h.data[j]
	return maxShift || minShift
}

func (h *Heap) isShiftBreak(i, j int) bool {
	maxShift := h.isMax && h.data[i] >= h.data[j]
	minShift := !h.isMax && h.data[i] <= h.data[j]
	return maxShift || minShift
}

func HeapSort(data []int, isMax bool) {
	n := len(data)
	h := NewHeap(n, isMax)

	for _, v := range data {
		h.Push(v)
	}

	for i := 0; i < n; i++ {
		data[i] = h.Pop()
	}
}

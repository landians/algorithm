package sorts

import (
	"container/heap"
)

type HeapArr []int

func (h *HeapArr) Len() int {
	return len(*h)
}

func (h *HeapArr) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *HeapArr) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *HeapArr) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *HeapArr) Pop() (v interface{}) {
	// 这里的操作可能是因为内存是按照小端的方式进行存储的
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func up(h *HeapArr, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(h *HeapArr, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

func (h *HeapArr) IsEmpty() bool {
	return h.Len() == 0
}

func (h *HeapArr) Sort() {
	heap.Init(h)

	heapSize := h.Len()

	heapSize--
	h.Swap(0, heapSize)

	for heapSize > 0 {
		if !down(h, 0, heapSize) {
			up(h, 0)
		}
		heapSize--
		h.Swap(0, heapSize)
	}
}

func sortedArrDistanceLessK(arr []int, k int) {
	h := HeapArr(make([]int, 0, len(arr)))

	min := len(arr)
	if min > k {
		min = k
	}

	index := 0
	for ; index < min; index++ {
		h.Push(arr[index])
	}

	i := 0
	for index < len(arr) {
		heap.Push(&h, arr[index])
		arr[i] = heap.Pop(&h).(int)
		i++
		index++
	}

	for !h.IsEmpty() {
		arr[i] = heap.Pop(&h).(int)
		i++
	}
}

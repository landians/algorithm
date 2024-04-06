package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
	data := []int{2, 1, 3, 6, 0, 4}
	maxHeap := NewHeapFrom(data, true)
	assert.Equal(t, []int{6, 3, 4, 1, 0, 2}, maxHeap.Data())

	minHeap := NewHeapFrom(data, false)
	assert.Equal(t, []int{0, 1, 3, 6, 2, 4}, minHeap.Data())
}

func TestHeapSort(t *testing.T) {
	d1 := []int{2, 1, 3, 6, 0, 4}
	HeapSort(d1, false)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 6}, d1)

	d2 := []int{2, 1, 3, 6, 0, 4}
	HeapSort(d2, true)
	assert.Equal(t, []int{6, 4, 3, 2, 1, 0}, d2)
}

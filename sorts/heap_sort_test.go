package sorts

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestHeapArr_Sort(t *testing.T) {
	heapArr := HeapArr{1, 2, 3, 4, 5, 6}
	heapArr.Sort()
	fmt.Println(heapArr)
}

func Test_HeapSort(t *testing.T) {
	heapArr := generateRandomArray(10000, 100000)
	arr1, arr2 := heapArr, heapArr

	sort.Ints(heapArr)
	comparator(arr2)

	assert.Equal(t, ordered(arr1), true, "arr1 应该有序")
	assert.Equal(t, arr1, arr2, "arr1 应该等于 arr2")
}

func Test_HeapSortedArrDistanceLessK(t *testing.T) {
	arr := []int{1, 5, 2, 3, 6, 4}
	sortedArrDistanceLessK(arr, 3)
	fmt.Println(arr)
	//assert.Equal(t, ordered(arr), true, "arr 应该有序")
}
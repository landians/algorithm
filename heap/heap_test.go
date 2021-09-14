package heap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_HeapSort(t *testing.T) {
	arr := generateRandomArray(10000, 100000)
	arr1, arr2 := arr, arr

	heapSort(arr1)
	comparator(arr2)

	assert.Equal(t, ordered(arr1), true, "arr1 应该有序")
	assert.Equal(t, arr1, arr2, "arr1 应该等于 arr2")
}

func Test_HeapInsert(t *testing.T) {
	fmt.Println("===== 测试 heapInsert =====")
	arr := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	fmt.Println(arr)
	fmt.Println("===== 测试 heapInsert =====")
}

func Test_HeapIfy(t *testing.T) {
	fmt.Println("===== 测试 heapIfy =====")
	arr := []int{1, 2, 3, 4, 5, 6}
	for i := len(arr)-1; i >=0; i-- {
		heapIfy(arr, i, len(arr))
	}
	fmt.Println(arr)
	fmt.Println("===== 测试 heapIfy =====")
}
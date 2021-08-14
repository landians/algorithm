package sorts

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initArray(maxSize, maxValue int) []int {
	return generateRandomArray(maxSize, maxValue)
}

func Test_GenerateRandomarray(t *testing.T) {
	fmt.Println("===== 随机数组生成 =====")
	arr := generateRandomArray(1000, 1000)
	fmt.Println(arr)
	fmt.Println("===== 随机数组生成 =====")
}

func Test_Equal(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{1, 3, 4, 5, 6, 7}
	assert.NotEqual(t, arr1, arr2, "arr1 应该不等于 arr2")
	arr3 := []int{1, 3, 4, 5, 6}
	assert.NotEqual(t, arr1, arr3, "arr1 应该不等于 arr3")
	arr4 := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, arr1, arr4, "arr1 应该等于 arr4")
}

func Test_Compartor(t *testing.T) {
	arr := generateRandomArray(100, 1000)
	comparator(arr)
	assert.Equal(t, ordered(arr), true, "arr 应该有序")
}

func Test_Ordered(t *testing.T) {
	arr := []int{2, 3, 1, -1, 10}
	assert.Equal(t, ordered(arr), false, "arr 应该无序")
	comparator(arr)
	assert.Equal(t, ordered(arr), true, "arr 应该有序")
}

func Test_SelectSort(t *testing.T) {
	arr := generateRandomArray(100000, 100000)
	arr1, arr2 := arr, arr

	SelectSort(arr1)
	comparator(arr2)

	assert.Equal(t, ordered(arr1), true, "arr1 应该有序")
	assert.Equal(t, arr1, arr2, "arr1 应该等于 arr2")
}

func Test_BubbleSort(t *testing.T) {
	arr := generateRandomArray(100000, 100000)
	arr1, arr2 := arr, arr

	BubbleSort(arr1)
	comparator(arr2)

	assert.Equal(t, ordered(arr1), true, "arr1 应该有序")
	assert.Equal(t, arr1, arr2, "arr1 应该等于 arr2")
}

func Test_InsertionSort(t *testing.T) {
	arr := generateRandomArray(100000, 100000)
	arr1, arr2 := arr, arr

	InsertionSort(arr)
	comparator(arr2)

	assert.Equal(t, ordered(arr1), true, "arr1 应该有序")
	assert.Equal(t, arr1, arr2, "arr1 应该等于 arr2")
}

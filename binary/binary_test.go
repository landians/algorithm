package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BinarySearch_1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	i := BinarySearch_1(arr, 3)
	assert.Equal(t, i, 2, "i 应该等于 2")

	i = BinarySearch_1(arr, 0)
	assert.Equal(t, i, 0, "i 应该等于 0")

	i = BinarySearch_1(arr, 7)
	assert.Equal(t, i, 5, "i 应该等于 5")
}

func Test_BinarySearch_2(t *testing.T) {
	arr := []int{6, 5, 4, 3, 2, 1}
	i := BinarySearch_2(arr, 3)
	assert.Equal(t, i, 3, "i 应该等于 3")

	i = BinarySearch_2(arr, 0)
	assert.Equal(t, i, 5, "i 应该等于 5")

	i = BinarySearch_2(arr, 7)
	assert.Equal(t, i, 0, "i 应该等于 5")
}

func Test_findLeftBound(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 5, 7}
	i := findLeftBound(nums, 3) // 查找到了从左往右数第一个 3
	assert.Equal(t, 2, i)
}

func Test_findRightBound(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 5, 7}
	i := findRightBound(nums, 3) // 查找到了从右往左数第一个 3
	assert.Equal(t, 5, i)
}

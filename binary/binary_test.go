package binary

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 2, 2, 3, 4, 4, 5}
	i := removeDuplicates(nums)
	assert.Equal(t, 5, i)
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	section := twoSum(nums, 9)
	assert.Equal(t, []int{1, 2}, section)

	nums = []int{2, 3, 4}
	section = twoSum(nums, 6)
	assert.Equal(t, []int{1, 3}, section)

	nums = []int{-1, 0}
	section = twoSum(nums, -1)
	assert.Equal(t, []int{1, 2}, section)
}

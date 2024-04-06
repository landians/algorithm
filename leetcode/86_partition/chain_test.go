package _6_partition

import (
	"github.com/landians/algorithm/leetcode"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition_1(t *testing.T) {
	l := leetcode.NewListFromArray([]int{1, 4, 3, 2, 5, 2})
	l1 := partition(l, 3)
	formatL1 := leetcode.FormatChain(l1)
	assert.Equal(t, "head->1->2->2->4->3->5", formatL1)
}

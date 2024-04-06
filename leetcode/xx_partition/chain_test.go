package xx_partition

import (
	"github.com/landians/algorithm/leetcode"
	"testing"
)

func TestPartition_1(t *testing.T) {
	l := leetcode.NewListFromArray([]int{1, 4, 3, 2, 5, 2})
	l1 := partition(l, 3)
	t.Log(leetcode.FormatChain(l1))
}

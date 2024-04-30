package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition_1(t *testing.T) {
	l := NewListFromArray([]int{1, 4, 3, 2, 5, 2})
	l1 := partition(l, 3)
	formatL1 := FormatChain(l1)
	assert.Equal(t, "head->1->2->2->4->3->5", formatL1)
}

func TestPartition_2(t *testing.T) {
	l := NewListFromArray([]int{1, 4, 3, 2, 5, 2})
	l1 := partition(l, 3)
	t.Log(FormatChain(l1))
}

func TestMergeTwoList(t *testing.T) {
	l1 := NewListFromArray([]int{1, 2, 3, 4, 5})
	l2 := NewListFromArray([]int{1, 2, 3, 4, 5})
	l := mergeTwoLists(l1, l2)
	formatL := FormatChain(l)
	assert.Equal(t, "head->1->1->2->2->3->3->4->4->5->5", formatL)
}

func TestIsPalindrome(t *testing.T) {
	l := NewListFromArray([]int{1, 2, 3, 2, 1})
	t.Log(isPalindrome(l))
}

package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLeftBound(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	i := FindLeftBound(nums, 8)
	assert.Equal(t, 3, i)
}

func TestFindRightBound(t *testing.T) {
	nums := []int{2, 2}
	i := FindLeftBound(nums, 3)
	t.Log(i)
}

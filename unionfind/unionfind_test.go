package unionfind

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestUnionSet(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	union := New(values)

	isSame := union.IsSameSet(1, 2)
	assert.Equal(t, false, isSame)

	isSame = union.IsSameSet(1, 3)
	assert.Equal(t, false, isSame)

	union.Union(1, 2)
	isSame = union.IsSameSet(1, 2)
	assert.Equal(t, true, isSame)

	isSame = union.IsSameSet(2, 3)
	assert.Equal(t, false, isSame)
}
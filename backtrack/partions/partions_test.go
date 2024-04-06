package partions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CanPartitionKSubsets(t *testing.T) {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	assert.Equal(t, true, canPartitionKSubsets(nums, 4))

	nums = []int{2, 2, 2, 2, 3, 4, 5}
	assert.Equal(t, false, canPartitionKSubsets(nums, 4))
}

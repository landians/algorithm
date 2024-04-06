package point24

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_JudgePoint24(t *testing.T) {
	nums := []int{4, 1, 8, 7}
	flag := judgePoint24(nums)
	assert.Equal(t, true, flag)

	nums = []int{1, 2, 1, 2}
	flag = judgePoint24(nums)
	assert.Equal(t, false, flag)
}
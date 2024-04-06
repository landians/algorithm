package coins

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	n := coinChange(coins, 11)
	assert.Equal(t, 3, n)

	coins = []int{2}
	n = coinChange(coins, 3)
	assert.Equal(t, -1, n)

	coins = []int{1}
	n = coinChange(coins, 0)
	assert.Equal(t, 0, n)

	coins = []int{1}
	n = coinChange(coins, 1)
	assert.Equal(t, 1, n)

	coins = []int{1}
	n = coinChange(coins, 2)
	assert.Equal(t, 2, n)

	coins = []int{186, 419, 83, 408}
	n = coinChange(coins, 6249)
	assert.Equal(t, 20, n)
}

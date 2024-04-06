package queen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolveNQueens(t *testing.T) {
	solution := solveNQueens(4)
	for _, v := range solution {
		t.Log(v)
	}
}

func Test_IsValid(t *testing.T) {
	board := initBoard(4)

	board[0][0] = "Q"

	assert.Equal(t, true, isValid(board, 2, 1))
}
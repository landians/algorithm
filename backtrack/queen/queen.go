package queen

import (
	"strings"
)

// N 皇后问题: https://leetcode-cn.com/problems/n-queens/

func solveNQueens(n int) [][]string {
	// N 皇后的摆放方法集合
	solution := make([][]string, 0)
	// '.' 表示空，'Q' 表示皇后，初始化空棋盘。
	board := initBoard(n)
	// 开始回溯摆放, 从第 0 行开始
	backtrack(board, 0, &solution)
	return solution
}

func initBoard(n int) [][]string {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	return board
}

// 路径：board 中小于 row 的那些行都已经成功放置了皇后
// 选择列表：第 row 行的所有列都是放置皇后的选择
// 结束条件：row 超过 board 的最后一行
func backtrack(board [][]string, row int, solution *[][]string) {
	if len(board) == row {
		rowSolution := make([]string, 0, row)
		for i := 0; i < row; i++ {
			rowContent := strings.Join(board[i], "")
			rowSolution = append(rowSolution, rowContent)
		}
		*solution = append(*solution, rowSolution)
		return
	}

	// 棋盘的列数获取
	n := len(board[row])

	// 尝试在每一列进行放置皇后
	for col := 0; col < n; col++ {
		// 排除不合法选择
		if !isValid(board, row, col) {
			continue
		}
		// 做选择
		board[row][col] = "Q"
		// 进入下一行决策
		backtrack(board, row+1, solution)
		// 撤销选择
		board[row][col] = "."
	}
}

func isValid(board [][]string, row int, col int) bool {
	// 获取棋盘的行数
	n := len(board)

	// 检查列是否有皇后冲突; 列号定死, 依次检测每一行
	for i := 0; i < n; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}

	// 检查右上方是否有皇后冲突; col ++， row--; 小于 0 为边界情况
	for colup, rowup := col+1, row-1; colup < n && rowup >= 0; colup, rowup = colup+1, rowup-1 {
		if board[rowup][colup] == "Q" {
			return false
		}
	}

	// 检查左上方是否有皇后冲突; col -- , row --; 小于 0 为边界情况
	for colup, rowup := col-1, row-1; colup >= 0 && rowup >= 0; colup, rowup = colup-1, rowup-1 {
		if board[rowup][colup] == "Q" {
			return false
		}
	}

	return true
}

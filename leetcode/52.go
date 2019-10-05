package leetcode

const (
	queen = 'Q'
	null  = '.'
)

func canPlace(board [][]byte, row, col int) bool {
	n := len(board)
	// 行
	for i := 0; i < col; i++ {
		if board[row][i] == queen {
			return false
		}
	}
	// 左斜上
	for i := 1; row-i >= 0 && col-i >= 0; i++ {
		if board[row-i][col-i] == queen {
			return false
		}
	}
	// 左斜下
	for i := 1; row+i < n && col-i >= 0; i++ {
		if board[row+i][col-i] == queen {
			return false
		}
	}
	return true
}

func nQueensRecur(board [][]byte, col int, result *int) {
	if col == len(board) {
		*result++
		return
	}
	for i := 0; i < len(board); i++ {
		if canPlace(board, i, col) {
			board[i][col] = queen
			nQueensRecur(board, col+1, result)
			board[i][col] = null
		}
	}
}

func totalNQueens(n int) int {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = null
		}
	}
	result := 0
	nQueensRecur(board, 0, &result)
	return result
}

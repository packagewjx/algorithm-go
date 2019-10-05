//+build 51

package leetcode

const (
	queen = 'Q'
	null  = '.'
)

func toStringBoard(board [][]byte) []string {
	res := make([]string, len(board))
	for i := 0; i < len(board); i++ {
		res[i] = string(board[i])
	}
	return res
}

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

func nQueensRecur(board [][]byte, col int, result *[][]string) {
	if col == len(board) {
		*result = append(*result, toStringBoard(board))
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

func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = null
		}
	}

	result := make([][]string, 0, n*n)
	nQueensRecur(board, 0, &result)
	return result
}

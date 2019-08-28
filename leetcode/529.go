package leetcode

func recursivelyReveal(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[x]) || board[x][y] == 'M' || board[x][y] == 'B' ||
		(board[x][y] >= '1' && board[x][y] <= '9') {
		return
	}

	mineCount := 0
	// 这里应该是E
	// 上一行
	for i := y - 1; x-1 >= 0 && i < len(board[x]) && i <= y+1; i++ {
		if i == -1 {
			continue
		}
		if board[x-1][i] == 'M' {
			mineCount++
		}
	}
	// 同一行
	if y-1 >= 0 && board[x][y-1] == 'M' {
		mineCount++
	}
	if y+1 < len(board[x]) && board[x][y+1] == 'M' {
		mineCount++
	}
	// 下一行
	for i := y - 1; x+1 < len(board) && i < len(board[x]) && i <= y+1; i++ {
		if i == -1 {
			continue
		}
		if board[x+1][i] == 'M' {
			mineCount++
		}
	}

	if mineCount > 0 {
		board[x][y] = byte('0' + mineCount)
		return
	}
	board[x][y] = 'B'
	recursivelyReveal(board, x-1, y-1)
	recursivelyReveal(board, x-1, y)
	recursivelyReveal(board, x-1, y+1)
	recursivelyReveal(board, x+1, y-1)
	recursivelyReveal(board, x+1, y)
	recursivelyReveal(board, x+1, y+1)
	recursivelyReveal(board, x, y-1)
	recursivelyReveal(board, x, y+1)
}

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	recursivelyReveal(board, click[0], click[1])
	return board
}

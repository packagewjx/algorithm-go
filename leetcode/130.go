package leetcode

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	var dfsCheck func(x, y int)
	dfsCheck = func(x, y int) {
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
			return
		}
		if visited[x][y] {
			return
		}
		visited[x][y] = true
		if board[x][y] == 'X' {
			return
		}
		// 走完所有的格子，设置全部visited状态
		dfsCheck(x-1, y)
		dfsCheck(x+1, y)
		dfsCheck(x, y-1)
		dfsCheck(x, y+1)
	}

	// 上边与下边
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			dfsCheck(0, i)
		}
		if board[len(board)-1][i] == 'O' {
			dfsCheck(len(board)-1, i)
		}
	}
	// 左边与右边
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			dfsCheck(i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			dfsCheck(i, len(board[0])-1)
		}
	}

	// 将没有visited的所有O改为X
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

package contest

func dominoRecur(placed int, i, j int, board [][]byte) int {
	if i == len(board)-1 && j == len(board[i])-1 {
		// 到达最后一个格子，返回
		return placed
	}
	nextI := i
	nextJ := j + 1
	if nextJ == len(board[i]) {
		nextJ = 0
		nextI += 1
	}
	if board[i][j] != Null {
		return dominoRecur(placed, nextI, nextJ, board)
	}
	max := placed
	// 尝试横着放
	if j+1 < len(board[i]) && board[i][j+1] == Null {
		board[i][j] = Full
		board[i][j+1] = Full
		result := dominoRecur(placed+1, nextI, nextJ, board)
		if result > max {
			max = result
		}
		board[i][j] = Null
		board[i][j+1] = Null
	}
	// 尝试竖着放
	if i+1 < len(board) && board[i+1][j] == Null {
		board[i+1][j] = Full
		board[i][j] = Full
		result := dominoRecur(placed+1, nextI, nextJ, board)
		if result > max {
			max = result
		}
		board[i][j] = Null
		board[i+1][j] = Null
	}
	// 尝试不放
	noPutResult := dominoRecur(placed, nextI, nextJ, board)
	if noPutResult > max {
		max = noPutResult
	}
	return max
}

const (
	Null = iota
	Full
	Broken
)

func domino(n int, m int, broken [][]int) int {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, m)
	}
	for _, bro := range broken {
		board[bro[0]][bro[1]] = Broken
	}

	return dominoRecur(0, 0, 0, board)
}

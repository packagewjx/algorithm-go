package leetcode

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	board := make([][]byte, 8)
	for i := 0; i < len(board); i++ {
		board[i] = make([]byte, 8)
	}

	for i := 0; i < len(queens); i++ {
		board[queens[i][0]][queens[i][1]] = 1
	}

	result := make([][]int, 0, 8)
	x := king[0]
	y := king[1]
	// 左上
	for i := 1; x-i >= 0 && y-i >= 0; i++ {
		if board[x-i][y-i] == 1 {
			result = append(result, []int{x - i, y - i})
			break
		}
	}

	// 上方
	for i := 1; x-i >= 0; i++ {
		if board[x-i][y] == 1 {
			result = append(result, []int{x - i, y})
			break
		}
	}

	// 右上
	for i := 1; x-i >= 0 && y+i < 8; i++ {
		if board[x-i][y+i] == 1 {
			result = append(result, []int{x - i, y + i})
			break
		}
	}

	// 右边
	for i := 1; y+i < 8; i++ {
		if board[x][y+i] == 1 {
			result = append(result, []int{x, y + i})
			break
		}
	}

	// 右下
	for i := 1; x+i < 8 && y+i < 8; i++ {
		if board[x+i][y+i] == 1 {
			result = append(result, []int{x + i, y + i})
			break
		}
	}

	// 下方
	for i := 1; x+i < 8; i++ {
		if board[x+i][y] == 1 {
			result = append(result, []int{x + i, y})
			break
		}
	}

	// 左下
	for i := 1; x+i < 8 && y-i >= 0; i++ {
		if board[x+i][y-i] == 1 {
			result = append(result, []int{x + i, y - i})
			break
		}
	}

	// 左边
	for i := 1; y-i >= 0; i++ {
		if board[x][y-i] == 1 {
			result = append(result, []int{x, y - i})
			break
		}
	}

	return result
}

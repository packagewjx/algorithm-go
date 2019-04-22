package leetcode

func numRookCaptures(board [][]byte) int {
	x := 0
	y := 0

	for x = 0; x < len(board); x++ {
		for y = 0; y < len(board[x]); y++ {
			if board[x][y] == 'R' {
				goto foundR
			}
		}
	}

foundR:
	count := 0

	temp := 0

	// go up
	for temp = x - 1; temp >= 0 && board[temp][y] == '.'; temp-- {
	}
	if temp >= 0 && board[temp][y] == 'p' {
		count++
	}

	// go down
	for temp = x + 1; temp < len(board) && board[temp][y] == '.'; temp++ {
	}
	if temp < len(board) && board[temp][y] == 'p' {
		count++
	}

	// go left
	for temp = y - 1; temp >= 0 && board[x][temp] == '.'; temp-- {
	}
	if temp >= 0 && board[x][temp] == 'p' {
		count++
	}

	// go right
	for temp = y + 1; temp < len(board[x]) && board[x][temp] == '.'; temp++ {

	}
	if temp < len(board[x]) && board[x][temp] == 'p' {
		count++
	}

	return count
}

package leetcode

func searchWord(board [][]byte, word string, matched int, x, y int, visited [][]bool) bool {
	if matched == len(word) {
		return true
	}
	visited[x][y] = true
	if x-1 >= 0 && !visited[x-1][y] && board[x-1][y] == word[matched] {
		if searchWord(board, word, matched+1, x-1, y, visited) {
			return true
		}
	}
	if x+1 < len(board) && !visited[x+1][y] && board[x+1][y] == word[matched] {
		if searchWord(board, word, matched+1, x+1, y, visited) {
			return true
		}
	}
	if y+1 < len(board[x]) && !visited[x][y+1] && board[x][y+1] == word[matched] {
		if searchWord(board, word, matched+1, x, y+1, visited) {
			return true
		}
	}
	if y-1 >= 0 && !visited[x][y-1] && board[x][y-1] == word[matched] {
		if searchWord(board, word, matched+1, x, y-1, visited) {
			return true
		}
	}
	visited[x][y] = false
	return false
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if searchWord(board, word, 1, i, j, visited) {
					return true
				}
			}
		}
	}
	return false
}

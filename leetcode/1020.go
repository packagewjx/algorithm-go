package leetcode

func visit1(A [][]int, x, y int, visited [][]bool) {
	if x < 0 || y < 0 || x >= len(visited) || y >= len(visited[x]) {
		return
	}
	if A[x][y] == 0 || visited[x][y] {
		return
	}

	visited[x][y] = true
	visit1(A, x+1, y, visited)
	visit1(A, x, y+1, visited)
	visit1(A, x-1, y, visited)
	visit1(A, x, y-1, visited)
}

func numEnclaves(A [][]int) int {
	visited := make([][]bool, len(A))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(A[0]))
	}

	// 上边界
	for i := 0; i < len(A[0]); i++ {
		if A[0][i] == 1 {
			visit1(A, 0, i, visited)
		}
	}
	// 下边界
	for i := 0; i < len(A[len(A)-1]); i++ {
		if A[len(A)-1][i] == 1 {
			visit1(A, len(A)-1, i, visited)
		}
	}
	// 左边界
	for i := 1; i < len(A)-1; i++ {
		if A[i][0] == 1 {
			visit1(A, i, 0, visited)
		}
	}
	// 右边界
	for i := 1; i < len(A)-1; i++ {
		if A[i][len(A[i])-1] == 1 {
			visit1(A, i, len(A[i])-1, visited)
		}
	}

	count := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if !visited[i][j] && A[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

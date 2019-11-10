package leetcode

func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}

	for i := 0; i < len(indices); i++ {
		for j := 0; j < m; j++ {
			matrix[indices[i][0]][j]++
		}
		for j := 0; j < n; j++ {
			matrix[j][indices[i][1]]++
		}
	}

	odd := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j]&1 == 1 {
				odd++
			}
		}
	}

	return odd
}

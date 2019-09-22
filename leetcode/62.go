package leetcode

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	path := make([][]int, m)
	for i := 0; i < len(path); i++ {
		path[i] = make([]int, n)
	}

	// 初始化
	for i := 0; i < n-1; i++ {
		path[m-1][i] = 1
	}
	for i := 0; i < m-1; i++ {
		path[i][n-1] = 1
	}

	// 填表
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			path[i][j] = path[i+1][j] + path[i][j+1]
		}
	}
	return path[0][0]
}

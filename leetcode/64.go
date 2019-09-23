package leetcode

func minPathSum(grid [][]int) int {
	count := make([][]int, len(grid))
	for i := 0; i < len(count); i++ {
		count[i] = make([]int, len(grid[0]))
	}

	count[len(grid)-1][len(grid[0])-1] = grid[len(grid)-1][len(grid[0])-1]
	for i := len(grid[0]) - 2; i >= 0; i-- {
		count[len(grid)-1][i] = grid[len(grid)-1][i] + count[len(grid)-1][i+1]
	}
	for i := len(grid) - 2; i >= 0; i-- {
		count[i][len(grid[0])-1] = grid[i][len(grid[0])-1] + count[i+1][len(grid[0])-1]
	}

	for i := len(grid) - 2; i >= 0; i-- {
		for j := len(grid[0]) - 2; j >= 0; j-- {
			if count[i+1][j] < count[i][j+1] {
				count[i][j] = grid[i][j] + count[i+1][j]
			} else {
				count[i][j] = grid[i][j] + count[i][j+1]
			}
		}
	}

	return count[0][0]
}

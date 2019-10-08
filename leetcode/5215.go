package leetcode

func getMaximumGoldDP(grid [][]int, row, col int, visited [][]bool) int {
	max := 0
	visited[row][col] = true
	if grid[row][col] != 0 {
		// 检查四个方向
		if row-1 >= 0 && !visited[row-1][col] {
			max = getMaximumGoldDP(grid, row-1, col, visited)
		}
		if row+1 < len(grid) && !visited[row+1][col] {
			m := getMaximumGoldDP(grid, row+1, col, visited)
			if m > max {
				max = m
			}
		}
		if col-1 >= 0 && !visited[row][col-1] {
			m := getMaximumGoldDP(grid, row, col-1, visited)
			if m > max {
				max = m
			}
		}
		if col+1 < len(grid[row]) && !visited[row][col+1] {
			m := getMaximumGoldDP(grid, row, col+1, visited)
			if m > max {
				max = m
			}
		}
	}
	visited[row][col] = false
	return max + grid[row][col]
}

func getMaximumGold(grid [][]int) int {
	maximum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			visited := make([][]bool, len(grid))
			for i := 0; i < len(grid); i++ {
				visited[i] = make([]bool, len(grid[i]))
			}
			dp := getMaximumGoldDP(grid, i, j, visited)
			if dp > maximum {
				maximum = dp
			}
		}
	}
	return maximum
}

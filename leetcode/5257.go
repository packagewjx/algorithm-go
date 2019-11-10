package leetcode

func closedIsland(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	var isIsland func(x, y int) bool
	isIsland = func(x, y int) bool {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) {
			return false
		}

		if visited[x][y] || grid[x][y] == 1 {
			return true
		}
		visited[x][y] = true

		down := isIsland(x+1, y)
		right := isIsland(x, y+1)
		up := isIsland(x-1, y)
		left := isIsland(x, y-1)
		return down && right && up && left
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 && !visited[i][j] {
				if isIsland(i, j) {
					count++
				}
			}
		}
	}

	return count
}

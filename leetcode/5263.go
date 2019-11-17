package leetcode

func shiftGrid(grid [][]int, k int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return grid
	}
	total := len(grid) * len(grid[0])
	col := len(grid[0])
	getNewPos := func(x, y int) (int, int) {
		pos := (x*col + y + k) % total
		return pos / col, pos % col
	}

	newGrid := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		newGrid[i] = make([]int, col)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < col; j++ {
			x, y := getNewPos(i, j)
			newGrid[x][y] = grid[i][j]
		}
	}

	return newGrid
}

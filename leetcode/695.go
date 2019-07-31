package leetcode

func countAreaAndClear(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + countAreaAndClear(grid, i-1, j) + countAreaAndClear(grid, i+1, j) +
		countAreaAndClear(grid, i, j-1) + countAreaAndClear(grid, i, j+1)
}

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				area := countAreaAndClear(grid, i, j)
				if area > max {
					max = area
				}
			}
		}
	}
	return max
}

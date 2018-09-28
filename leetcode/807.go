package leetcode

func maxIncreaseKeepingSkyline(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	result := 0

	vertSkyline, horiSkyline := calVerSkylineAndHoriSkyline(grid)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			result += min(vertSkyline[y], horiSkyline[x]) - grid[x][y]
		}
	}

	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func calVerSkylineAndHoriSkyline(grid [][]int) (vertSkyline, horiSkyline []int) {
	vertSkyline = make([]int, len(grid))
	horiSkyline = make([]int, len(grid[0]))

	//y为纵轴，x为横轴
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] > horiSkyline[y] {
				horiSkyline[y] = grid[y][x]
			}
			if grid[y][x] > vertSkyline[x] {
				vertSkyline[x] = grid[y][x]
			}
		}
	}

	return
}

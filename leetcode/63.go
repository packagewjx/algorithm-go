package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 1 {
		for i := 0; i < len(obstacleGrid[0]); i++ {
			if obstacleGrid[0][i] == 1 {
				return 0
			}
		}
		return 1
	} else if len(obstacleGrid[0]) == 1 {
		for i := 0; i < len(obstacleGrid); i++ {
			if obstacleGrid[i][0] == 1 {
				return 0
			}
		}
		return 1
	} else if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}

	pathCount := make([][]int, len(obstacleGrid))
	for i := 0; i < len(pathCount); i++ {
		pathCount[i] = make([]int, len(obstacleGrid[0]))
	}
	pathCount[len(obstacleGrid)-1][len(obstacleGrid[0])-1] = 1

	for i := len(obstacleGrid) - 1; i >= 0; i-- {
		for j := len(obstacleGrid[0]) - 1; j >= 0; j-- {
			if (i == len(obstacleGrid)-1 && j == len(obstacleGrid[0])-1) || obstacleGrid[i][j] == 1 {
				continue
			}

			if i == len(obstacleGrid)-1 {
				pathCount[i][j] = pathCount[i][j+1]
			} else if j == len(obstacleGrid[0])-1 {
				pathCount[i][j] = pathCount[i+1][j]
			} else {
				pathCount[i][j] = pathCount[i+1][j] + pathCount[i][j+1]
			}
		}
	}
	return pathCount[0][0]
}

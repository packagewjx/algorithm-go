package leetcode

import "math"

func setZeroes(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				for k := 0; k < len(matrix); k++ {
					if matrix[k][j] == 0 {
						continue
					}
					matrix[k][j] = math.MinInt64
				}
				for k := 0; k < len(matrix[0]); k++ {
					if matrix[i][k] == 0 {
						continue
					}
					matrix[i][k] = math.MinInt64
				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == math.MinInt64 {
				matrix[i][j] = 0
			}
		}
	}
}

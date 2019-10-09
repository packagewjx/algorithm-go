//+build 85

package leetcode

func maximalRectangleDP(matrix [][]byte, row, col int, length int, memo [][][]int) int {
	if memo[row][col][length] != -1 {
		return memo[row][col][length]
	}

	area := 0
	// 实际计算
	// 首先查看从row col开始是否有长为length的线段

	if matrix[row][col] == '0' {
		area = 0
	} else /* matrix[row][col]=='1' */ {
		// 查看memo[row][col+1][length-1]是否大于0，来确认是否有长为length的线段。因为这个值如果大于0，则至少是length-1
		if length > 1 && memo[row][col+1][length-1] == 0 {
			area = 0
		} else {
			// 只有有这么长的线段才计算
			if row+1 < len(matrix) {
				// 如果有下一行，则查看下一行的这个长度是否有矩形
				nextArea := maximalRectangleDP(matrix, row+1, col, length, memo)
				area = nextArea + length
			} else {
				// 没有下一行的话，面积就是这个线段的长度
				area = length
			}
		}

	}

	memo[row][col][length] = area
	return area
}

func maximalRectangle(matrix [][]byte) int {
	memo := make([][][]int, len(matrix))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([][]int, len(matrix[i]))
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = make([]int, len(memo[i])+1)
			for k := 0; k < len(memo[i][j]); k++ {
				memo[i][j][k] = -1
			}
		}
	}

	max := 0
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			for k := 1; j+k <= len(matrix[i]); k++ {
				area := maximalRectangleDP(matrix, i, j, k, memo)
				if area == 0 {
					// area为0说明没有这么长的线段
					// 后面的也可以设置为0
					for l := k + 1; j+l <= len(matrix[i]); l++ {
						memo[i][j][l] = 0
					}
					break
				}
				if area > max {
					max = area
				}
			}
		}
	}

	return max
}

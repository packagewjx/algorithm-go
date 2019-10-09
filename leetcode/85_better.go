package leetcode

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	// widths[i][j]保存着到(i,j)的连续长方形的最大长度
	widths := make([][]int, len(matrix))
	for i := 0; i < len(widths); i++ {
		widths[i] = make([]int, len(matrix[i]))
		if matrix[i][0] == '1' {
			widths[i][0] = 1
		}
		for j := 1; j < len(widths[i]); j++ {
			if matrix[i][j] == '1' {
				widths[i][j] = widths[i][j-1] + 1
			} else {
				widths[i][j] = 0
			}
		}
	}

	// 使用栈解题目
	max := 0
	stack := make([]int, 1, len(matrix))
	stack[0] = -1
	for col := 0; col < len(widths[0]); col++ {
		for row := 0; row < len(widths); row++ {
			for len(stack) > 1 && widths[stack[len(stack)-1]][col] >= widths[row][col] {
				area := (row - stack[len(stack)-2] - 1) * widths[stack[len(stack)-1]][col]
				if area > max {
					max = area
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, row)
		}

		for len(stack) > 1 {
			area := (len(widths) - stack[len(stack)-2] - 1) * widths[stack[len(stack)-1]][col]
			if area > max {
				max = area
			}
			stack = stack[:len(stack)-1]
		}
	}

	return max
}

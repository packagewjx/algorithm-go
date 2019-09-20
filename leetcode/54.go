package leetcode

// 螺旋矩阵题，考验编程思想的题目
func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	height := len(matrix)
	width := len(matrix[0])
	res := make([]int, 0, height*width)
	up := 0
	down := height - 1
	right := width - 1
	left := 0
	for true {
		// 上面的
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		// 右边的
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		// 下面的
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if down < up {
			break
		}
		// 左边的
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return res
}

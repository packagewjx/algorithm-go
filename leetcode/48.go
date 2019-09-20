package leetcode

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < len(matrix)-i-1; j++ {
			// 上方的旋转90度
			nRow := j
			nCol := len(matrix) - 1 - i
			t1 := matrix[nRow][nCol]
			matrix[nRow][nCol] = matrix[i][j]
			// 右方的旋转90度
			nRow = len(matrix) - 1 - i
			nCol = len(matrix) - 1 - j
			t2 := matrix[nRow][nCol]
			matrix[nRow][nCol] = t1
			// 下方的旋转90度
			nRow = len(matrix) - 1 - j
			nCol = i
			t1 = matrix[nRow][nCol]
			matrix[nRow][nCol] = t2
			// 放回原本的位置
			matrix[i][j] = t1
		}
	}
}

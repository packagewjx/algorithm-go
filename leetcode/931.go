package leetcode

func minFallingPathSum(A [][]int) int {
	rowCnt := len(A)
	colCnt := len(A[0])
	min := make([][]int, rowCnt)
	for i := 0; i < rowCnt; i++ {
		min[i] = make([]int, colCnt)
	}

	// 最后一行
	last := A[rowCnt-1]
	for i := 0; i < len(last); i++ {
		min[rowCnt-1][i] = last[i]
	}

	for i := rowCnt - 2; i >= 0; i-- {
		for j := 0; j < colCnt; j++ {
			smallest := 1<<63 - 1
			for k := j - 1; k < j+2; k++ {
				if k < 0 || k >= colCnt {
					continue
				}
				if min[i+1][k] < smallest {
					smallest = min[i+1][k]
				}
			}
			min[i][j] = smallest + A[i][j]
		}
	}

	smallest := 1<<63 - 1
	for i := 0; i < colCnt; i++ {
		if smallest > min[0][i] {
			smallest = min[0][i]
		}
	}
	return smallest
}

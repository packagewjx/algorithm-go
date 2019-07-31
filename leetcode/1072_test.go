package leetcode

import "testing"

func Test1072(t *testing.T) {
	matrix := [][]int{{1, 1, 0, 0}, {0, 0, 1, 1}, {1, 0, 1, 0}, {0, 0, 1, 1}}
	println(maxEqualRowsAfterFlips(matrix))
}

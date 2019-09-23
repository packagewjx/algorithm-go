package leetcode

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	// 查看第一个数是否小于
	if matrix[0][0] > target {
		return false
	}
	// 先寻找行
	row := 0
	for ; row < len(matrix) && matrix[row][0] <= target; row++ {
	}
	row--
	pos := sort.SearchInts(matrix[row], target)
	return pos < len(matrix[row]) && matrix[row][pos] == target
}

package leetcode

import "testing"

func Test63(t *testing.T) {
	println(uniquePathsWithObstacles([][]int{
		{0, 0, 0, 1},
		{0, 1, 0, 0},
		{1, 0, 1, 0},
		{0, 0, 0, 0},
	}))
}

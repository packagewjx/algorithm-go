package leetcode

import "testing"

func Test547(t *testing.T) {
	println(findCircleNum([][]int{
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1},
	}))
}

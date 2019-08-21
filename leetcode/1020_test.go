package leetcode

import "testing"

func Test1020(t *testing.T) {
	println(numEnclaves([][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 1, 0, 0, 0}}))
}

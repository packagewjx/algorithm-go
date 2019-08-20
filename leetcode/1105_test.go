package leetcode

import "testing"

func Test1104(t *testing.T) {
	test := [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}
	println(minHeightShelves(test, 4))
}

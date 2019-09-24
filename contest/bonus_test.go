package contest

import "testing"

func TestBonus(t *testing.T) {
	//输入：N = 6, leadership = [[1, 2], [1, 6], [2, 3], [2, 5], [1, 4]], operations = [[1, 1, 500], [2, 2, 50], [3, 1], [2, 6, 15], [3, 1]]
	bonus(6, [][]int{{1, 2}, {1, 6}, {2, 3}, {2, 5}, {1, 4}}, [][]int{{1, 1, 500}, {2, 2, 50}, {3, 1}, {2, 6, 15}, {3, 1}})
}

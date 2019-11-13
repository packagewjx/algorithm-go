package leetcode

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	res := make([][]int, 2)
	res[0] = make([]int, len(colsum))
	res[1] = make([]int, len(colsum))
	for i := 0; i < len(colsum); i++ {
		if colsum[i] == 2 {
			res[0][i] = 1
			res[1][i] = 1
			upper--
			lower--
		}
	}

	// 贪心法
	for i := 0; i < len(colsum); i++ {
		if colsum[i] == 1 {
			// 变为0立即返回
			if upper == 0 && lower == 0 {
				return [][]int{}
			}
			if upper > lower {
				res[0][i] = 1
				upper--
			} else {
				res[1][i] = 1
				lower--
			}
		}
	}

	if upper == 0 && lower == 0 {
		return res
	} else {
		return [][]int{}
	}
}

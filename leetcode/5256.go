package leetcode

import "math/rand"

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	res := make([][]int, 2)
	res[0] = make([]int, len(colsum))
	res[1] = make([]int, len(colsum))
	// 填上2
	if colsum[len(colsum)-1] == 2 {
		res[0][len(colsum)-1] = 1
		res[1][len(colsum)-1] = 1
		upper--
		lower--
	}
	test := make([]int, 0, len(colsum)/2)
	for i := len(colsum) - 2; i >= 0; i-- {
		if colsum[i] == 2 {
			res[0][i] = 1
			res[1][i] = 1
			upper--
			lower--
		} else if colsum[i] == 1 {
			test = append(test, i)
		}
	}

	var bt func(pos int) bool
	bt = func(pos int) bool {
		if pos == len(test) {
			return true
		}

		r := rand.Intn(1)
		if r == 1 {
			// 尝试放上面
			if upper > 0 {
				res[0][test[pos]] = 1
				upper--
				if bt(pos + 1) {
					return true
				}
				res[0][test[pos]] = 0
				upper++
			}

			// 尝试放下面
			if lower > 0 {
				res[1][test[pos]] = 1
				lower--
				if bt(pos + 1) {
					return true
				}
				res[1][test[pos]] = 0
				lower++
			}
		} else {
			// 尝试放下面
			if lower > 0 {
				res[1][test[pos]] = 1
				lower--
				if bt(pos + 1) {
					return true
				}
				res[1][test[pos]] = 0
				lower++
			}
			// 尝试放上面
			if upper > 0 {
				res[0][test[pos]] = 1
				upper--
				if bt(pos + 1) {
					return true
				}
				res[0][test[pos]] = 0
				upper++
			}
		}

		return false
	}

	if bt(0) {
		return res
	} else {
		return nil
	}
}

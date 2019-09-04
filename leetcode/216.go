package leetcode

func combinationSum3Recur(k int, n int, used []bool, result *[][]int, cur []int, start int) {
	if k == 1 {
		if n < len(used) && !used[n] && n >= start {
			r := make([]int, len(cur)+1)
			copy(r, cur)
			r[len(r)-1] = n
			*result = append(*result, r)
		}
		return
	}
	for i := start; i <= 9 && i <= n; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		cur = append(cur, i)
		combinationSum3Recur(k-1, n-i, used, result, cur, i+1)
		// 复原
		used[i] = false
		cur = cur[:len(cur)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	used := make([]bool, 10)
	result := make([][]int, 0, 10)
	combinationSum3Recur(k, n, used, &result, []int{}, 1)
	return result
}

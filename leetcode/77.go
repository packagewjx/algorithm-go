package leetcode

func combineRecursive(cur []int, start, n, k int, result *[][]int) {
	if len(cur) == k {
		*result = append(*result, cur)
		return
	}

	// 剪枝，若剩余的数不够填满，则一定是没有结果的
	if k-len(cur) > n-start+1 {
		return
	}

	for i := start; i <= n; i++ {
		newChoice := make([]int, len(cur)+1, k)
		copy(newChoice, cur)
		newChoice[len(cur)] = i
		combineRecursive(newChoice, i+1, n, k, result)
	}
}

func combine(n int, k int) [][]int {
	if k > n {
		return [][]int{}
	} else if k == n {
		res := make([]int, n)
		for i := 0; i < n; i++ {
			res[i] = i + 1
		}
		return [][]int{res}
	}

	result := make([][]int, 0, 1<<uint(k+1))

	combineRecursive([]int{}, 1, n, k, &result)
	return result
}

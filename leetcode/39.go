package leetcode

func combinationSumBT(candidates []int, target int, cur []int) [][]int {
	result := make([][]int, 0, 4)
	for i := 0; i < len(candidates); i++ {
		if candidates[i] < target {
			n := make([]int, len(cur)+1)
			copy(n, cur)
			n[len(cur)] = candidates[i]
			// candidates[i:]代表能够再选择本数，但是不能再选择前面的数，可以去重
			bt := combinationSumBT(candidates[i:], target-candidates[i], n)
			result = append(result, bt...)
		} else if candidates[i] == target {
			n := make([]int, len(cur)+1)
			copy(n, cur)
			n[len(cur)] = candidates[i]
			result = append(result, n)
		}
	}
	return result
}

func combinationSum(candidates []int, target int) [][]int {
	return combinationSumBT(candidates, target, []int{})
}

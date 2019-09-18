package leetcode

import "sort"

func combinationSum2BT(candidates []int, target int, cur []int) [][]int {
	result := make([][]int, 0, 4)
	for i := 0; i < len(candidates); i++ {
		// 剪枝。如果前一个数和这个数相等，则可以继续，前一个数的集合更大，一定包含这个数的结果
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] == target {
			n := make([]int, len(cur)+1)
			copy(n, cur)
			n[len(cur)] = candidates[i]
			result = append(result, n)
		} else if candidates[i] < target {
			n := make([]int, len(cur)+1)
			copy(n, cur)
			n[len(cur)] = candidates[i]
			bt := combinationSum2BT(candidates[i+1:], target-candidates[i], n)
			result = append(result, bt...)
		}
	}
	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] <= candidates[j]
	})
	return combinationSum2BT(candidates, target, []int{})
}

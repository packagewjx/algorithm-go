package leetcode

func subsetsRecursive(element []int, start, remain int, selected []int, result *[][]int) {
	if start+remain == len(element) {
		// 没得选了，直接返回
		set := make([]int, len(selected), len(selected)+remain)
		copy(set, selected)
		set = append(set, element[start:]...)
		*result = append(*result, set)
		return
	} else if start+remain > len(element) {
		// 不可能的，返回
		return
	}

	// 依次选择
	for i := start; i+remain-1 < len(element); i++ {
		set := make([]int, len(selected), len(selected)+remain)
		copy(set, selected)
		set = append(set, element[i])
		if remain == 1 {
			//选完了加入结果
			*result = append(*result, set)
		} else {
			// 否则继续
			subsetsRecursive(element, i+1, remain-1, set, result)
		}
	}
}

func subsets(nums []int) [][]int {
	result := make([][]int, 2, 1<<uint(len(nums)))
	result[0] = []int{}
	result[1] = nums
	for i := 1; i < len(nums); i++ {
		subsetsRecursive(nums, 0, i, []int{}, &result)
	}
	return result
}

package leetcode

import "sort"

func subsetsWithDupBT(nums []int, pos int, memo [][][]int) [][]int {
	if pos == len(nums) {
		return [][]int{{}}
	} else if memo[pos] != nil {
		return memo[pos]
	}

	result := make([][]int, 0, 1<<uint(len(nums)-pos))
	// 查看pos的数是否重复
	numEnd := pos + 1
	for ; numEnd < len(nums) && nums[numEnd] == nums[pos]; numEnd++ {

	}

	bt := subsetsWithDupBT(nums, numEnd, memo)
	// 依次带入0到numEnd-pos个数到集合中
	temp := make([]int, 0, numEnd-pos)
	for i := pos; i <= numEnd; i++ {
		for j := 0; j < len(bt); j++ {
			newResult := make([]int, len(temp), len(nums))
			copy(newResult, temp)
			newResult = append(newResult, bt[j]...)
			result = append(result, newResult)
		}
		// 添加一个数到temp中，代表选择了这个数
		temp = append(temp, nums[pos])
	}

	memo[pos] = result
	return result
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	memo := make([][][]int, len(nums))
	r := subsetsWithDupBT(nums, 0, memo)
	return r
}

package leetcode

import "sort"

func reverseSlice47(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		temp := slice[i]
		slice[i] = slice[len(slice)-1-i]
		slice[len(slice)-1-i] = temp
	}
}

// 更一般化的下一个排列算法，支持非自然数序列，支持重复元素
func nextPermutation47(nums []int) bool {
	pos := len(nums) - 1
	// 从后面开始找往前面升序的序列，找到第一个元素
	for ; pos > 0 && nums[pos] <= nums[pos-1]; pos-- {
	}
	// 这里是判断是否结束的条件。pos为0代表从最后一直到最前面都是升序的
	if pos == 0 {
		return false
	}
	pos = pos - 1
	// 寻找pos后面最后的一个大于nums[pos]的数
	lastBigger := len(nums) - 1
	for ; nums[lastBigger] <= nums[pos] && lastBigger > pos; lastBigger-- {
	}
	if lastBigger == pos {
		// 没有找到更大的数字，是不寻常的，起码nums[pos+1]大
		panic("impossible")
	}
	// 交换
	temp := nums[pos]
	nums[pos] = nums[lastBigger]
	nums[lastBigger] = temp

	// 将后面的slice反转
	reverseSlice47(nums[pos+1:])
	return true
}

// 一般化全排列算法，可包含重复数字
func permuteUnique(nums []int) [][]int {
	// 为了防止重复，首先排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})

	res := make([][]int, 0, 16)
	first := make([]int, len(nums))
	res = append(res, first)
	copy(first, nums)
	hasNext := nextPermutation47(nums)
	for hasNext {
		r := make([]int, len(nums))
		copy(r, nums)
		res = append(res, r)
		hasNext = nextPermutation47(nums)
	}
	return res
}

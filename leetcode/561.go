package leetcode

import "sort"

func arrayPairSum(nums []int) int {
	return doArrayPariSum1(nums)
}

func doArrayPariSum1(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

//TODO 还有更快的解法
func doArrayPariSum2(nums []int) int {
	return 0
}

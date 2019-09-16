package leetcode

import (
	"math"
	"sort"
)

func threeSumTarget(nums []int, target int) [][]int {
	lastNum := math.MinInt64
	result := make([][]int, 0, 10)
	for i := 0; i < len(nums)-2; i++ {
		goal := target - nums[i]
		if lastNum == nums[i] {
			// 若已经访问过相同的数字则跳过
			continue
		}
		lastNum = nums[i]
		// 固定i为最左边的数，而从其后方选取，这样避免重复
		p := i + 1
		q := len(nums) - 1
		for p < q {
			if p > i+1 && nums[p] == nums[p-1] {
				// 跳过i和相等的数字
				p++
				continue
			} else if q < len(nums)-1 && nums[q] == nums[q+1] {
				q--
				continue
			}
			s := nums[p] + nums[q]
			if s > goal {
				q--
			} else if s < goal {
				p++
			} else /* s == goal */ {
				result = append(result, []int{nums[i], nums[p], nums[q]})
				p++
				q--
			}
		}
	}
	return result
}

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0, 10)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i]-nums[j] < 0
	})

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		goal := target - nums[i]
		sumTarget := threeSumTarget(nums[i+1:], goal)
		if len(sumTarget) > 0 {
			for j := 0; j < len(sumTarget); j++ {
				result = append(result, append(sumTarget[j], nums[i]))
			}
		}
	}
	return result
}

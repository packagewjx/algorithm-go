package leetcode

import (
	"math"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i]-nums[j] < 0
	})

	lastNum := math.MinInt64
	result := make([][]int, 0, 10)
	for i := 0; i < len(nums)-2; i++ {
		target := -nums[i]
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
			if s > target {
				q--
			} else if s < target {
				p++
			} else /* s == target */ {
				result = append(result, []int{nums[i], nums[p], nums[q]})
				p++
				q--
			}
		}
	}
	return result
}

package leetcode

import "sort"

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	used := make(map[int]int)

	// 构造有序数组
	type Number struct {
		val   int
		label int
	}
	nums := make([]*Number, len(values))
	for i := 0; i < len(values); i++ {
		nums[i] = &Number{
			val:   values[i],
			label: labels[i],
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i].val > nums[j].val
	})

	sum := 0
	totalUsed := 0
	// 贪婪法
	for i := 0; totalUsed < num_wanted && i < len(nums); i++ {
		if used[nums[i].label] < use_limit {
			sum += nums[i].val
			used[nums[i].label] += 1
			totalUsed++
		}
	}
	return sum
}

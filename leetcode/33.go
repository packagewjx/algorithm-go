package leetcode

import "sort"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if nums[0] < nums[len(nums)-1] {
		// 这个数组没有旋转
		pos := sort.SearchInts(nums, target)
		if pos < len(nums) && nums[pos] == target {
			return pos
		} else {
			return -1
		}
	}

	// 查看是在左半部分还是右半部分
	if nums[0] <= target {
		// 左半部分
		start := nums[0]
		begin := 0
		end := len(nums)
		for begin < end {
			mid := (begin + end) / 2
			if nums[mid] == target {
				return mid
			} else if nums[mid] < target && nums[mid] > start {
				// 这个数在左半边的数组中，或者他小于
				begin = mid + 1
			} else {
				end = mid
			}
		}
	} else if nums[len(nums)-1] >= target {
		// 右半部分
		last := nums[len(nums)-1]

		begin := 0
		end := len(nums)
		for begin < end {
			mid := (begin + end) / 2
			if nums[mid] == target {
				return mid
			} else if nums[mid] > last || nums[mid] < target {
				// 这个数在左半边的数组中，或者他小于
				begin = mid + 1
			} else {
				end = mid
			}
		}
	}

	// 到这里就是没有找到
	return -1
}

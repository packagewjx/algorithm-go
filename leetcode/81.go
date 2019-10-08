package leetcode

import "sort"

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	// 排除与左右两边相等的情况
	if target == nums[0] || target == nums[len(nums)-1] {
		return true
	}
	// 需要去除两端的重复数字，使得两个不相同
	l := 0
	r := len(nums)
	if nums[0] == nums[len(nums)-1] {
		l = 1
		for ; l < len(nums) && nums[l] == nums[0]; l++ {
		}
		r = len(nums) - 1
		for ; r > 0 && nums[r-1] == nums[len(nums)-1]; r-- {
		}
	}
	if l > r {
		// 说明整个数组都是一个数字
		return false
	}

	minLeft := nums[l]
	maxRight := nums[r-1]
	if target == minLeft || target == maxRight {
		return true
	} else if minLeft < maxRight {
		// 数组没有旋转
		pos := sort.SearchInts(nums, target)
		return pos < len(nums) && nums[pos] == target
	} else {
		if target > minLeft {
			// 目标在左边
			for l < r {
				m := l + (r-l)/2
				if nums[m] == target {
					return true
				} else if nums[m] < minLeft || nums[m] > target {
					r = m
				} else {
					l = m + 1
				}
			}
		} else {
			// 目标在右边
			for l < r {
				m := l + (r-l)/2
				if nums[m] == target {
					return true
				} else if nums[m] > maxRight || nums[m] < target {
					l = m + 1
				} else {
					r = m
				}
			}
		}
		return l < len(nums) && nums[l] == target
	}
}

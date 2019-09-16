package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	closestSum := 0
	closestVariation := math.MaxInt64

	sort.Slice(nums, func(i, j int) bool {
		return nums[i]-nums[j] < 0
	})

	for i := 0; i < len(nums)-2; i++ {
		goal := target - nums[i]
		p := i + 1
		q := len(nums) - 1

		for p < q {
			sum := nums[p] + nums[q]
			v := goal - sum
			if v < 0 {
				v = -v
			} else if v == 0 {
				// 能达到，直接返回
				return target
			}
			// 更新最接近和
			if v < closestVariation {
				closestSum = sum + nums[i]
				closestVariation = v
			}
			// 调整v
			if sum < goal {
				p++
			} else /* sum > goal */ {
				q--
			}

		}
	}
	return closestSum
}

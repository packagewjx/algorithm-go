package leetcode

import "math"

func jump(nums []int) int {
	step := make([]int, 0, len(nums))
	for i := 0; i < len(step); i++ {
		step[i] = math.MaxInt64
	}

	return step[len(nums)-1]
}

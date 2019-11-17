package leetcode

import (
	"math"
)

func maxSumDivThree(nums []int) int {
	minMod := [3][2]int{{math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64}}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		mod := nums[i] % 3
		switch mod {
		case 0:
			continue
		default:

			if nums[i] < minMod[mod][0] {
				temp := minMod[mod][0]
				minMod[mod][0] = nums[i]
				if temp < minMod[mod][1] {
					minMod[mod][1] = temp
				}
			} else if nums[i] < minMod[mod][1] {
				minMod[mod][1] = nums[i]
			}
		}
	}

	mod := sum % 3

	switch mod {
	case 2:
		if minMod[2][0] == math.MaxInt64 {
			// 肯定有两个 mod3为1的数
			return sum - minMod[1][0] - minMod[1][1]
		} else {
			// 存在mod3为2的数
			if minMod[1][0] != math.MaxInt64 && minMod[1][1] != math.MaxInt64 && minMod[1][0]+minMod[1][1] < minMod[2][0] {
				return sum - minMod[1][0] - minMod[1][1]
			} else {
				return sum - minMod[2][0]
			}
		}
	case 1:
		if minMod[1][0] == math.MaxInt64 {
			// 肯定存在两个mod3为2的数
			return sum - minMod[2][0] - minMod[2][1]
		} else {
			if minMod[2][0] != math.MaxInt64 && minMod[2][1] != math.MaxInt64 && minMod[2][0]+minMod[2][1] < minMod[1][0] {
				return sum - minMod[2][0] - minMod[2][1]
			} else {
				return sum - minMod[1][0]
			}
		}
	case 0:
		return sum
	default:
		panic("impossble")
	}
}

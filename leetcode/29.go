package leetcode

import (
	"math"
	"sort"
)

func divide(dividend int, divisor int) int {
	// 处理边界情况
	if divisor == 1 {
		return dividend
	} else if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		} else {
			return -dividend
		}
	}

	negative := false
	if dividend < 0 && divisor > 0 {
		negative = true
		dividend = -dividend
	} else if dividend > 0 && divisor < 0 {
		negative = true
		divisor = -divisor
	} else if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	}
	if dividend < divisor {
		return 0
	}

	result := 1
	// 二次演进
	multiply := make([]int, 0, 16)
	cur := divisor
	for cur <= dividend {
		multiply = append(multiply, cur)
		cur *= 2
		result <<= 1
	}
	// 得到了目前最大的商，cur就是余数
	cur = dividend - multiply[len(multiply)-1]
	result >>= 1
	// 余数继续在multiply数组中寻找最大的
	for cur >= multiply[0] {
		pos := sort.SearchInts(multiply, cur)
		// 刚好除得尽
		if multiply[pos] == cur {
			result += 1 << uint(pos)
			break
		}
		// 除不尽的话，这个pos一定是大于0的，因为cur >= multiply[0]
		pos--
		result += 1 << uint(pos)
		cur -= multiply[pos]
	}

	if negative {
		return -result
	} else {
		return result
	}
}

package leetcode

import "strconv"

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}

	// 第一个数除以后面所有数顺序除的商，因为后面所有数相除，要么会得到小于1的最小的小数，要么得到大于1的最小的整数，两者都是最大的结果
	result := strconv.Itoa(nums[0]) + "/("
	for i := 1; i < len(nums); i++ {
		result += strconv.Itoa(nums[i]) + "/"
	}
	result = result[:len(result)-1]
	result += ")"
	return result
}

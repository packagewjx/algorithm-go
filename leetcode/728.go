package leetcode

func selfDividingNumbers(left int, right int) []int {
	nums := make([]int, 0, (right-left)/2)
	for i := left; i <= right; i++ {
		if isSelfDivide(i) {
			nums = append(nums, i)
		}
	}
	return nums
}

func isSelfDivide(num int) bool {
	originalNum := num
	for num > 0 {
		digit := num % 10
		num = num / 10
		if digit == 0 {
			// skip 0
			return false
		}

		if originalNum%digit != 0 {
			return false
		}
	}
	return true
}

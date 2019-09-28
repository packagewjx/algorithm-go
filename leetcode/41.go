package leetcode

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > len(nums) || (i != nums[i]-1 && nums[nums[i]-1] == nums[i]) {
			// 当原本位置上已经有合适的数，或者数超过数组，或者数小于等于0的时候，设置为-1
			nums[i] = -1
		} else if nums[i] != i+1 {
			temp := nums[i]
			nums[i] = nums[temp-1]
			nums[temp-1] = temp
			i--
		}
	}

	start := 0
	for start = 0; start < len(nums) && nums[start] != -1; start++ {

	}

	return start + 1
}

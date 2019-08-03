package leetcode

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	begin := 0
	end := len(nums)

	for begin+1 < end {
		pos := begin + (end-begin)/2
		// 保证pos是单数，这样在单独数左边的都是num[pos]==num[pos-1]，右边的都是num[pos]!=num[pos-1]
		pos += 1 - (pos & 1)
		if nums[pos] == nums[pos-1] {
			begin = pos + 1
		} else {
			end = pos
		}
	}

	return nums[begin]
}

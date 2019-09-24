package leetcode

func sortColors(nums []int) {
	zero := 0
	two := len(nums) - 1
	swap := func(i, j int) {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}

	for i := 0; i <= two; {
		if nums[i] == 0 {
			if zero == i {
				i++
			} else {
				swap(zero, i)
			}
			zero++
		} else if nums[i] == 2 {
			swap(two, i)
			two--
		} else {
			i++
		}
	}
}

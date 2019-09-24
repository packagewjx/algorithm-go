package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	num := nums[0] - 1
	end := 0
	count := 0
	swap := func(i, j int) {
		if i == j {
			return
		}
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}
	for i := 0; i < len(nums); i++ {
		if num != nums[i] {
			num = nums[i]
			count = 1
			swap(i, end)
			end++
		} else if count < 2 {
			swap(i, end)
			end++
			count++
		}
	}
	return end
}

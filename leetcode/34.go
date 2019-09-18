package leetcode

// 返回等于或者第一个大于等于的数
func binarySearchLeftMost(nums []int, target int) int {
	begin := 0
	end := len(nums)
	for begin < end {
		mid := begin + (end-begin)/2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			begin = mid + 1
		} else /* nums[mid] == target */ {
			if mid == 0 {
				// 第一个是的话，肯定是最左了
				return 0
			}
			// 如果前一个不是target的话可以直接返回了。如果前一个是target，则可以排除本数
			if nums[mid-1] != target {
				return mid
			} else {
				end = mid
			}
		}
	}
	return begin
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	i1 := binarySearchLeftMost(nums, target)
	if i1 == len(nums) || nums[i1] != target {
		// 不存在这个数，直接返回
		return []int{-1, -1}
	}
	i2 := binarySearchLeftMost(nums, target+1)
	return []int{i1, i2 - 1}
}

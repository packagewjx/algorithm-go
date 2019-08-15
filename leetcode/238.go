package leetcode

// 返回从最后一个数乘过来的值
func productExceptSelfRecursive(cur int, nums []int, result []int) int {
	if cur == len(nums)-1 {
		return nums[cur]
	}
	result[cur+1] = result[cur] * nums[cur]
	latterMultiply := productExceptSelfRecursive(cur+1, nums, result)
	result[cur] = latterMultiply * result[cur]
	return latterMultiply * nums[cur]
}

// 递归算法，进入第i层递归的时候，意味着检查数组第i个数，此时result中存储着从0到i-1的数的乘积，然后将nums[i]乘以这个乘积，存放到result[i+1]
// 供下一层递归使用。当进入到最后的一层时，这里的result就是从前面乘过来的数字，无需改变，而返回这一层的数字，代表从后面开始乘过来的结果。返回到
// 第i层时，就得到从第i+1到后面乘过来的结果，此时将这个数字与result[i]相乘，便是除去i数字的乘积。
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = 1
	productExceptSelfRecursive(0, nums, result)
	return result
}

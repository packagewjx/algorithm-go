package leetcode

func reverseSlice(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		temp := slice[i]
		slice[i] = slice[len(slice)-1-i]
		slice[len(slice)-1-i] = temp
	}
}

/**
生成permute的下一个字典序排列。返回是否还有下一个，如果有，permute将会是新的下一个，如果没有，则permute不变
*/
func nextPermutation(permute []int) bool {
	pos := len(permute) - 1
	for ; pos >= 1 && permute[pos] < permute[pos-1]; pos-- {
	}
	// 这里是判断是否结束的条件
	if pos == 0 {
		return false
	}
	pos = pos - 1
	// 寻找最后的一个大于permute[pos]的数
	lastBigger := len(permute) - 1
	for ; permute[lastBigger] < permute[pos]; lastBigger-- {
	}
	// 交换
	{
		temp := permute[pos]
		permute[pos] = permute[lastBigger]
		permute[lastBigger] = temp
	}

	// 将后面的slice反转
	reverseSlice(permute[pos+1:])

	return true
}

/**
使用经典的字典序排列算法
*/
func permute(nums []int) [][]int {
	permutation := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		permutation[i] = i
	}

	ret := make([][]int, 0, 10)
	hasNext := true
	for hasNext {
		result := make([]int, len(nums))
		for i := 0; i < len(permutation); i++ {
			result[i] = nums[permutation[i]]
		}
		ret = append(ret, result)
		hasNext = nextPermutation(permutation)
	}
	return ret
}

package leetcode

func canJump(nums []int) bool {
	// lastPos代表能够跳到最后的点的最后一个位置
	lastPos := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

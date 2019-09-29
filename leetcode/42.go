package leetcode

func max42(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func leftRightMaxDP(height []int, pos int, leftMax []int, rightMax []int) {
	if pos == 0 {
		leftMax[pos] = 0
		leftRightMaxDP(height, 1, leftMax, rightMax)
		// 返回时，rightMax应该是设定好了
		rightMax[0] = max42(rightMax[1], height[1])
	} else if pos == len(height)-1 {
		rightMax[pos] = 0
		leftMax[pos] = max42(leftMax[pos-1], height[pos-1])
	} else {
		leftMax[pos] = max42(leftMax[pos-1], height[pos-1])
		leftRightMaxDP(height, pos+1, leftMax, rightMax)
		rightMax[pos] = max42(rightMax[pos+1], height[pos+1])
	}
}

// 使用动态规划得出当前位置左边和右边的最大高度，然后再计算处水的高度
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftRightMaxDP(height, 0, leftMax, rightMax)

	res := 0
	for i := 0; i < len(height); i++ {
		if rightMax[i] <= leftMax[i] && rightMax[i] > height[i] {
			res += rightMax[i] - height[i]
		} else if rightMax[i] >= leftMax[i] && leftMax[i] > height[i] {
			res += leftMax[i] - height[i]
		}
	}
	return res
}

// 双指针法
func trapDoublePointer(height []int) int {
	l := 0
	r := len(height) - 1

	res := 0
	leftMax := 0
	rightMax := 0
	for l < r {
		if height[l] < height[r] {
			if height[l] >= leftMax {
				leftMax = height[l]
			} else {
				/*
					如果左指针的高度低过右指针，意味着当前水的高度由leftMax决定。
					因为左边最高的时候是leftMax，意味着这个地方高度小于leftMax。当我们在leftMax位置的时候，进入的是上一个分支，
					意味着左边肯定，是比右边某个位置低了，可能是rightMax的位置，也可能是某个比rightMax更高的位置
					因此在后面的所有这个位置的遍历，都是leftMax决定水的高度。
				*/
				res += leftMax - height[l]
			}
			l++
		} else {
			if height[r] >= rightMax {
				rightMax = height[r]
			} else {
				res += rightMax - height[r]
			}
			r--
		}
	}
	return res
}

package leetcode

func longestOnes(A []int, K int) int {
	if K == 0 {
		max := 0
		currentStart := -1
		currentLength := 0
		for i := 0; i < len(A); i++ {
			if A[i] == 1 {
				if currentStart == -1 {
					currentStart = i
					currentLength = 1
				} else {
					currentLength++
				}
			} else /* A[i] == 0 */ {
				if currentLength > max {
					max = currentLength
				}
				currentStart = -1
				currentLength = 0
			}
		}
		if currentLength > max {
			max = currentLength
		}
		return max
	}

	// 保存我填的1的队列
	queue := make([]int, 0, K)
	// 首先填入3个1
	window := 0
	// 填满开始的空位
	for i := 0; i < len(A) && len(queue) < cap(queue); i++ {
		if A[i] == 0 {
			queue = append(queue, i)
		}
		window++
	}
	if len(queue) < cap(queue) {
		return window
	}

	max := window
	// 从最后1个1的后一个开始
	for i := queue[len(queue)-1] + 1; i < len(A); i++ {
		if A[i] == 0 {
			// 此时window必定是最大的，因为要从前面取1来填到这里
			if window > max {
				max = window
			}
			onePos := queue[0]
			// 放1到新的位置
			queue = queue[1:]
			queue = append(queue, i)
			// 计算window的减少值
			// 若最前面是我们填的1，此时窗口是不变的，只有不是的时候才处理
			if i-window != onePos {
				window -= onePos - (i - window)
			}

		} else /* A[i] == 1 */ {
			window++
		}
	}
	if window > max {
		max = window
	}
	return max
}

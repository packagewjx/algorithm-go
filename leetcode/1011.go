package leetcode

// 装袋问题
func shipWithinDays(weights []int, D int) int {
	begin := 0
	end := 0
	for i := 0; i < len(weights); i++ {
		if weights[i] > begin {
			begin = weights[i]
		}
		end += weights[i]
	}

	for begin < end {
		capacity := begin + (end-begin)/2
		cur := 0
		count := 0
		for i := 0; i < len(weights); i++ {
			if cur+weights[i] > capacity {
				count++
				cur = 0
			}
			cur += weights[i]
		}
		count++
		if count > D {
			begin = capacity + 1
		} else {
			end = capacity
		}
	}
	return end
}

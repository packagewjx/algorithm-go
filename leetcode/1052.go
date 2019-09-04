package leetcode

func maxSatisfied(customers []int, grumpy []int, X int) int {
	sum := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] != 1 {
			sum += customers[i]
		}
	}

	// 初始的窗口
	max := sum
	for i := 0; i < X; i++ {
		if grumpy[i] == 1 {
			max += customers[i]
		}
	}
	cur := max
	for i := X; i < len(customers); i++ {
		if grumpy[i] == 1 {
			cur += customers[i]
		}
		if grumpy[i-X] == 1 {
			cur -= customers[i-X]
		}
		if cur > max {
			max = cur
		}
	}
	return max
}

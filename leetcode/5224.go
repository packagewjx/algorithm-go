package leetcode

func dieSimBT(rest int, rollMax []int, cur int, selected [7]int, leftMax [7]int, countPermute func(selected [7]int) int64) int64 {
	if cur == 6 {
		if rest > rollMax[5] {
			return 0
		} else {
			selected[cur] = rest
			res := countPermute(selected)
			selected[cur] = 0
			return res
		}
	} else if rest == 0 {
		// 计算结果
		return countPermute(selected)
	}
	// 选择当前骰子0次到rollMax次
	result := int64(0)
	for i := 0; i <= rollMax[cur-1] && i <= rest; i++ {
		selected[cur] = i
		result += dieSimBT(rest-i, rollMax, cur+1, selected, leftMax, countPermute)
	}

	selected[cur] = 0
	return result
}

func dieSimulator(n int, rollMax []int) int {
	perm := make([]int64, n+1)
	perm[0] = 1
	for i := 1; i <= n; i++ {
		perm[i] = int64(i) * perm[i-1]
	}
	var leftMax [7]int
	leftMax[6] = rollMax[5]
	for i := 5; i > 0; i-- {
		leftMax[i] = rollMax[i-1] + leftMax[i+1]
	}

	countPermute := func(selected [7]int) int64 {
		total := perm[n]
		for i := 1; i <= 6; i++ {
			if selected[i] > 1 {
				total /= perm[selected[i]]
			}
		}
		return total
	}

	return int(dieSimBT(n, rollMax, 1, [7]int{}, leftMax, countPermute) % 1000000007)
}

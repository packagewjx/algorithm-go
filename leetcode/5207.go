package leetcode

func equalSubstring(s string, t string, maxCost int) int {
	cost := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] > t[i] {
			cost[i] = int(s[i] - t[i])
		} else {
			cost[i] = int(t[i] - s[i])
		}
	}

	maxWindow := 0
	winLen := 0
	curCost := 0
	for i := 0; i < len(s); i++ {
		if curCost+cost[i] <= maxCost {
			curCost += cost[i]
			winLen++
		} else {
			// 减去前面的值，直到可以加入这个值
			for j := i - winLen; j < i; j++ {
				curCost -= cost[j]
				winLen--
				if curCost+cost[i] <= maxCost {
					break
				}
			}
			if curCost+cost[i] <= maxCost {
				curCost += cost[i]
				winLen++
			}
		}

		if winLen > maxWindow {
			maxWindow = winLen
		}
	}
	return maxWindow
}

package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([]int, len(prices)+1)
	highest := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > highest {
			highest = prices[i]
			dp[i] = dp[i+1]
		} else {
			p := highest - prices[i]
			if p > dp[i+1] {
				dp[i] = p
			} else {
				dp[i] = dp[i+1]
			}
		}
	}

	profit := 0
	lowest := 0x7fffffff
	// 首先考察0位置的交易，作为初始最大交易
	twoMax := dp[0]

	for i := 0; i < len(prices); i++ {
		if prices[i] < lowest {
			lowest = prices[i]
		}
		thisProfit := prices[i] - lowest
		if thisProfit > profit {
			profit = thisProfit
		}

		// 查看前面最大的profit与i+1位置的和
		newTwoMax := profit + dp[i+1]
		if newTwoMax > twoMax {
			twoMax = newTwoMax
		}
	}

	return twoMax
}

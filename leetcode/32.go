package leetcode

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}

	startPos := make([]int, 0, len(s)>>1)
	// 结束位置为key，开始位置为value的map
	validPosMap := make(map[int]int)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			startPos = append(startPos, i)
		} else /*s[i] == ')'*/ {
			if len(startPos) > 0 {
				// 说明有的配对
				start := startPos[len(startPos)-1]
				startPos = startPos[:len(startPos)-1]
				// 查看是否可以接上前面的有效对
				p, ok := validPosMap[start]
				if ok {
					// 有可以接上的
					delete(validPosMap, start)
					validPosMap[i+1] = p
				} else {
					// 没有可接上的。i+1是因为结束位置是不包含的
					validPosMap[i+1] = start
				}
			}
		}
	}

	// 在validPosMap找最长的
	longest := 0
	for endPos, startPos := range validPosMap {
		l := endPos - startPos
		if l > longest {
			longest = l
		}
	}
	return longest
}

// 动态规划法
func longestValidParenthesesDP(s string) int {
	// 定义dp为每个位置的最长有效括号长度
	dp := make([]int, len(s))
	max := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else {
			if i > 0 {
				if s[i-1] == '(' {
					// 当前i位置可以和前一个配对，形如`....()`，那么最长长度，应该是2加上i-2位置的长度
					if i-2 >= 0 {
						dp[i] = 2 + dp[i-2]
					} else {
						dp[i] = 2
					}
				} else /*s[i-1]=')'*/ {
					// 当前是`...))`，那么，如果i-1的位置的子字符串，其前一个位置是`(`，这就配对了i的`)`
					// 记得还需要加上与这个i配对的括号的前一个
					if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
						dp[i] = dp[i-1] + 2
						if i-dp[i-1]-2 >= 0 {
							dp[i] += dp[i-dp[i-1]-2]
						}
					}
				}
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
	}
	return max
}

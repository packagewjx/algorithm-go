package leetcode

import "math"

func minCut(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 0
	dp[len(s)-1] = 0

	for i := len(s) - 2; i >= 0; i-- {
		min := math.MaxInt64

		for mid := i; mid+(mid-i) < len(s); mid++ {
			// 以j为中心的回文串
			isHuiwen := true
			for j := 0; mid-j >= i; j++ {
				if s[mid-j] != s[mid+j] {
					isHuiwen = false
					break
				}
			}
			if isHuiwen {
				res := 0
				// 这条回文串是整条字符串，mid处在中间，则无需分割，结果为0。否则，需要+1
				if (len(s)+i)/2 != mid {
					res = 1 + dp[mid+(mid-i)+1]
				}
				if res < min {
					min = res
				}
			}

			// 以j和j+1为中心的回文串
			if mid+mid-i+1 < len(s) && s[mid] == s[mid+1] {
				isHuiwen = true
				for j := 1; mid-j >= i; j++ {
					if s[mid-j] != s[mid+1+j] {
						isHuiwen = false
						break
					}
				}
				if isHuiwen {
					res := 0
					if (len(s)+i-1)/2 != mid {
						res = 1 + dp[mid+(mid-i+1)+1]
					}
					if res < min {
						min = res
					}
				}
			}
		}
		dp[i] = min
	}

	return dp[0]
}

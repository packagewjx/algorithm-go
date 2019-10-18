package leetcode

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 0; i <= len(s); i++ {
		dp[i][0] = 1
	}

	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			// 我不选这个字母结束，我看看s的i-1个字符串中有多少个t中j个的子序列
			dp[i][j] += dp[i-1][j]
			// 如果s[i-1]与t[j-1]相等的话，我可以选上这个字母，再看看s的i-1字母中有多少个t中j-1个字母的子序列
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[len(s)][len(t)]
}

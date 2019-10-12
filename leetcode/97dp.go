package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	// 二维动态规划，根据官方题解编程
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else if s1[i-1] == s3[i+j-1] {
				// 貌似可以让s1的某长度子串放到这里s3这里结尾的位置，需要进一步查看
				dp[i][j] = dp[i-1][j]
			} else if s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = false
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

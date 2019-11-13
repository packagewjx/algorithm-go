//+build 139

package leetcode

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return s == ""
	}

	wordMap := make(map[string]bool)
	longest := 0
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = true
		if len(wordDict[i]) > longest {
			longest = len(wordDict[i])
		}
	}

	dp := make([]bool, len(s)+1)
	dp[len(s)] = true
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j <= len(s) && j-i <= longest; j++ {
			if wordMap[s[i:j]] {
				if dp[j] {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[0]

}

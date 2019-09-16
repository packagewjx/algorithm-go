package leetcode

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	longest := ""
	for i := 0; i < len(s); i++ {
		// 检查从i为中心的
		for j := 1; i-j >= 0 && i+j < len(s) && s[i-j] == s[i+j]; j++ {
			t := s[i-j : i+j+1]
			if len(t) > len(longest) {
				longest = t
			}
		}
		// 检查以i和i+1为中心的
		if i+1 < len(s) {
			for j := 0; i-j >= 0 && i+1+j < len(s) && s[i-j] == s[i+1+j]; j++ {
				t := s[i-j : i+j+2]
				if len(t) > len(longest) {
					longest = t
				}
			}
		}
	}
	if len(longest) == 0 {
		longest = s[0:1]
	}
	return longest
}

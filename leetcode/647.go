package leetcode

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		// 检查从这个位置开始的最长的单数长度回文字符串
		length := 1
		for j := 1; i-j >= 0 && i+j < len(s); j++ {
			if s[i-j] == s[i+j] {
				length += 2
			} else {
				break
			}
		}
		// 若最长的长度为2n+1，则有n+1条子回文串
		count += length/2 + 1

		// 检查从这个位置开始的最长的双数长度回文字符串
		length = 0
		for j := 1; i-j+1 >= 0 && i+j < len(s); j++ {
			if s[i-j+1] == s[i+j] {
				length += 2
			} else {
				break
			}
		}
		// 最长长度为2n的话，则有n条回文子字符串
		count += length / 2
	}
	return count
}

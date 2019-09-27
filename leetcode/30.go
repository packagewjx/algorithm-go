package leetcode

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	targetCount := make([]int, 27)
	wordLen := len(words[0])
	totalLen := len(words) * wordLen
	wordCount := make(map[string]int)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			targetCount[words[i][j]&31]++
		}
		wordCount[words[i]]++
	}

	if totalLen > len(s) {
		return []int{}
	}

	// 计算初始totalLen的情况
	count := make([]int, 27)
	for i := 0; i < totalLen; i++ {
		count[s[i]&31]++
	}
	// matched代表这个统计表中，count与targetCount相同的部分
	matched := 0
	for i := 1; i < 27; i++ {
		if targetCount[i] == count[i] {
			matched++
		}
	}

	isMatch := func(str string) bool {
		count := make(map[string]int)
		for i := 0; i < len(str); i += wordLen {
			matchWord := str[i : i+wordLen]
			count[matchWord]++
			if count[matchWord] > wordCount[matchWord] {
				return false
			}
		}
		return true
	}

	result := make([]int, 0)
	for i := 0; i+totalLen < len(s); i++ {
		if matched == 26 {
			// 若是26，代表本子字符串与words的有同样数量的字符，可以查看是否是组成的
			if isMatch(s[i : i+totalLen]) {
				result = append(result, i)
			}
		}

		// 准备加入下一个字符，减去当前的字符

		key := s[i] & 31
		count[key]--
		if count[key] == targetCount[key] {
			matched++
		} else if count[key]+1 == targetCount[key] {
			matched--
		}

		key = s[i+totalLen] & 31
		count[key]++
		if count[key] == targetCount[key] {
			matched++
		} else if count[key]-1 == targetCount[key] {
			matched--
		}
	}
	// 比对最后一个子串
	if matched == 26 {
		if isMatch(s[len(s)-totalLen:]) {
			result = append(result, len(s)-totalLen)
		}
	}
	return result
}

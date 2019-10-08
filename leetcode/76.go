package leetcode

func minWindow(s string, t string) string {
	if s == "" || t == "" || len(s) < len(t) {
		return ""
	}

	targetCount := make([]int, 128)
	for i := 0; i < len(t); i++ {
		targetCount[t[i]]++
	}

	distinctT := make(map[byte]bool)
	for i := 'a'; i <= 'z'; i++ {
		if targetCount[i] > 0 {
			distinctT[byte(i)] = true
		}
	}
	for i := 'A'; i <= 'Z'; i++ {
		if targetCount[i] > 0 {
			distinctT[byte(i)] = true
		}
	}

	// 保证开始的第一个是在字符集的，不加入额外的无用字符，保证最短字串
	start := 0
	for ; start < len(s) && !distinctT[s[start]]; start++ {
	}
	// 如果start已经到尾部，说明没有
	if start == len(s) {
		return ""
	}

	charPos := make([]int, 0, 50)
	winMatched := 0
	winWord := make([]int, 128)
	shortest := s

	for i := start; i < len(s); i++ {
		char := s[i]
		if distinctT[char] {
			// 存在，则加1
			winWord[char]++
			charPos = append(charPos, i)
			if winWord[char] == targetCount[char] {
				// 原本不等的，现在加了之后相等，match加1
				winMatched++
			}

			// 去除开头的多余的字符
			for winWord[s[charPos[0]]] > targetCount[s[charPos[0]]] {
				winWord[s[charPos[0]]]--
				charPos = charPos[1:]
			}

			if winMatched == len(distinctT) {
				if i-charPos[0]+1 < len(shortest) {
					shortest = s[charPos[0] : i+1]
				}
			}
		}
	}
	if winMatched == len(distinctT) {
		return shortest
	} else {
		return ""
	}
}

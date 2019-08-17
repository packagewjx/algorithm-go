package study

func lcsRecursive(s1, s2 string, s1Start, s2Start int, memo [][]map[string]bool) map[string]bool {
	if memo[s1Start][s2Start] != nil {
		return memo[s1Start][s2Start]
	}

	longestLength := 0
	lcs := map[string]bool{"": true}

	handleSubResult := func(subResult map[string]bool) {
		subResultLength := 0
		// 寻找其中一个，查看其长度
		for sequence, _ := range subResult {
			subResultLength = len(sequence)
		}

		if subResultLength > longestLength {
			// 复制
			lcs = make(map[string]bool)
			for sequences, _ := range subResult {
				lcs[sequences] = true
			}
			longestLength = subResultLength
		} else if subResultLength == longestLength {
			for sequence, _ := range subResult {
				lcs[sequence] = true
			}
		}
	}

	for i := len(s1) - 1; i > s1Start; i-- {
		for j := len(s2) - 1; j > s2Start; j-- {
			subResult := lcsRecursive(s1, s2, i, j, memo)
			// 这个subResult起码有一个结果
			handleSubResult(subResult)
		}
	}
	// 判断本位置
	if s1[s1Start] == s2[s2Start] {
		newLcs := make(map[string]bool)
		for sequence, _ := range lcs {
			newLcs[s1[s1Start:s1Start+1]+sequence] = true
		}
		lcs = newLcs
	} else {
		if s2Start+1 < len(s2) {
			subResult := lcsRecursive(s1, s2, s1Start, s2Start+1, memo)
			handleSubResult(subResult)
		}
		if s1Start+1 < len(s1) {
			subResult := lcsRecursive(s1, s2, s1Start+1, s2Start, memo)
			handleSubResult(subResult)
		}

	}

	memo[s1Start][s2Start] = lcs
	return lcs
}

func LongestCommonSequence(s1, s2 string) []string {
	memo := make([][]map[string]bool, len(s1))
	for i := 0; i < len(s1); i++ {
		memo[i] = make([]map[string]bool, len(s2))
	}
	// 初始化最后一格
	if s1[len(s1)-1] == s2[len(s2)-1] {
		memo[len(s1)-1][len(s2)-1] = map[string]bool{s1[len(s1)-1:]: true}
	} else {
		memo[len(s1)-1][len(s2)-1] = map[string]bool{"": true}
	}

	m := lcsRecursive(s1, s2, 0, 0, memo)
	result := make([]string, 0, len(m))
	for sequence, _ := range m {
		result = append(result, sequence)
	}
	return result
}

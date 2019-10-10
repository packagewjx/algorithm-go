package leetcode

func isScrambleDC(s1 string, s1Sum []int, s1SumBase int, s2 string, s2Sum []int, s2SumBase int) bool {
	if s1 == s2 {
		return true
	} else if len(s1) != len(s2) {
		return false
	} else if len(s1) == 1 {
		return s1 == s2
	}

	getSum := func(prefixSum []int, base int, start, end int) int {
		if start == 0 {
			return prefixSum[end-1] - base
		} else {
			return prefixSum[end-1] - prefixSum[start-1]
		}
	}
	isSameAlphabet := func(str1, str2 string) bool {
		if len(str1) != len(str2) {
			return false
		}

		c1 := make([]int, 27)
		c2 := make([]int, 27)
		for i := 0; i < len(str1); i++ {
			c1[str1[i]&31]++
			c2[str2[i]&31]++
		}
		for i := 1; i < len(c1); i++ {
			if c1[i] != c2[i] {
				return false
			}
		}
		return true
	}

	result := false
	// 使用加法作为哈希，快速查看两条字符串是否有相同的字符集。和不同时能够准确判断没有相同字符集，但是如果和相同，需要进一步测试
	l1Sum := 0
	l2Sum := 0
	r2Sum := 0
	for i := 1; i < len(s1); i++ {
		l1Sum = getSum(s1Sum, s1SumBase, 0, i)
		l2Sum = getSum(s2Sum, s2SumBase, 0, i)
		r2Sum = getSum(s2Sum, s2SumBase, len(s2)-i, len(s2))
		if l1Sum == l2Sum {
			rest1Sum := getSum(s1Sum, s1SumBase, i, len(s1))
			rest2Sum := getSum(s2Sum, s2SumBase, i, len(s2))
			if rest1Sum == rest2Sum {
				// 可能在i点分割了
				if isSameAlphabet(s1[:i], s2[:i]) && isSameAlphabet(s1[i:], s2[i:]) {
					// 分治，检验是否正确
					if isScrambleDC(s1[:i], s1Sum[:i], s1SumBase, s2[:i], s2Sum[:i], s2SumBase) &&
						isScrambleDC(s1[i:], s1Sum[i:], s1Sum[i-1], s2[i:], s2Sum[i:], s2Sum[i-1]) {
						result = true
						break
					}
				}
			}
		}
		if l1Sum == r2Sum {
			// s1
			rest1Sum := getSum(s1Sum, s1SumBase, i, len(s1))
			rest2Sum := getSum(s2Sum, s2SumBase, 0, len(s2)-i)
			if rest1Sum == rest2Sum {
				// 可能在i点分割，并且交换了
				if isSameAlphabet(s1[:i], s2[len(s2)-i:]) && isSameAlphabet(s1[i:], s2[:len(s2)-i]) {
					// 分治，检验是否正确
					if isScrambleDC(s1[:i], s1Sum[:i], s1SumBase, s2[len(s2)-i:], s2Sum[len(s2)-i:], s2Sum[len(s2)-i-1]) &&
						isScrambleDC(s1[i:], s1Sum[i:], s1Sum[i-1], s2[:len(s2)-i], s2Sum[:len(s2)-i], s2SumBase) {
						result = true
						break
					}
				}
			}
		}
	}
	return result
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1PrefixSum := make([]int, len(s1))
	s2PrefixSum := make([]int, len(s2))
	s1PrefixSum[0] = int(s1[0])
	s2PrefixSum[0] = int(s2[0])
	for i := 1; i < len(s1); i++ {
		s1PrefixSum[i] = s1PrefixSum[i-1] + int(s1[i])
		s2PrefixSum[i] = s2PrefixSum[i-1] + int(s2[i])
	}

	return isScrambleDC(s1, s1PrefixSum, 0, s2, s2PrefixSum, 0)
}

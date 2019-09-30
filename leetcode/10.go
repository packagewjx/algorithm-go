//+build leetcode10

package leetcode

func isMatchDP(s string, p string, si, pi int, memo [][]*bool) bool {
	if memo[si][pi] != nil {
		return *memo[si][pi]
	}
	isMatch := false

	if len(p)-pi >= 2 && p[pi+1] == '*' {
		matchingEnd := si
		if p[pi] == '.' {
			matchingEnd = len(s)
		} else {
			for ; matchingEnd < len(s) && s[matchingEnd] == p[pi]; matchingEnd++ {
			}
		}
		// 使用s[si:matchingEnd]的字符，消费掉这个"${p[pi]}*"
		for i := si; i <= matchingEnd; i++ {
			if isMatchDP(s, p, i, pi+2, memo) {
				isMatch = true
				break
			}
		}
	} else {
		if pi < len(p) && si < len(s) && (p[pi] == '.' || s[si] == p[pi]) {
			isMatch = isMatchDP(s, p, si+1, pi+1, memo)
		} else {
			isMatch = false
		}
	}

	memo[si][pi] = &isMatch
	return isMatch
}

// 回溯算法
func isMatch(s string, p string) bool {
	memo := make([][]*bool, len(s)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]*bool, len(p)+1)
	}

	for i := 0; i < len(s); i++ {
		memo[i][len(p)] = new(bool)
	}
	t := true
	memo[len(s)][len(p)] = &t
	return isMatchDP(s, p, 0, 0, memo)
}

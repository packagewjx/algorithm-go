//+build 44dp

package leetcode

const (
	NIL   = 0
	TRUE  = 1
	FALSE = -1
)

func isMatchDP(s string, p string, sPos, pPos int, memo [][]int) int {
	if memo[sPos][pPos] != NIL {
		return memo[sPos][pPos]
	}

	// 下面的语句有pPos < len(p)
	result := NIL
	if p[pPos] == '*' {
		// 用0个到多个匹配这个*
		for i := sPos; i <= len(s); i++ {
			if isMatchDP(s, p, i, pPos+1, memo) == TRUE {
				result = TRUE
				break
			}
		}
		if result == NIL {
			result = FALSE
		}
	} else if sPos < len(s) && (s[sPos] == p[pPos] || p[pPos] == '?') {
		result = isMatchDP(s, p, sPos+1, pPos+1, memo)
	} else {
		result = FALSE
	}

	memo[sPos][pPos] = result
	return result
}

func isMatch(s string, p string) bool {
	memo := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		memo[i] = make([]int, len(p)+1)
		memo[i][len(p)] = FALSE
	}
	memo[len(s)][len(p)] = TRUE
	return isMatchDP(s, p, 0, 0, memo) == TRUE
}

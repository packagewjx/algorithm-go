package leetcode

func numDecodingsBT(s string, pos int, memo []int) int {
	if pos == len(s) {
		return 1
	} else if memo[pos] != -1 {
		return memo[pos]
	} else if s[pos] == '0' {
		memo[pos] = 0
		return 0
	}

	// 非0的开始

	result := 0
	// 1个数解码
	code := s[pos] & 15
	result += numDecodingsBT(s, pos+1, memo)
	if len(s)-1 != pos {
		code = code*10 + s[pos+1]&15
		if code <= 26 {
			result += numDecodingsBT(s, pos+2, memo)
		}
	}

	memo[pos] = result
	return result
}

func numDecodings(s string) int {
	// memo[i]为从第i个开始的解码方式的总数
	memo := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		memo[i] = -1
	}

	return numDecodingsBT(s, 0, memo)
}

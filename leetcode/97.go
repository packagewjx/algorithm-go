//+build 97

package leetcode

const (
	NIL97   = 0
	TRUE97  = 1
	FALSE97 = 2
)

func isInterleaveDP(s1, s2, s3 string, s1Pos, s2Pos, s3Pos int, memo [][][]byte) byte {
	if memo[s1Pos][s2Pos][s3Pos] != NIL97 {
		return memo[s1Pos][s2Pos][s3Pos]
	} else if s3Pos == len(s3) {
		return TRUE97
	} else if s1Pos == len(s1) {
		// 此时s1已用完，查看剩余部分是否等于s2
		if s2[s2Pos:] == s3[s3Pos:] {
			return TRUE97
		} else {
			return FALSE97
		}
	} else if s2Pos == len(s2) {
		if s1[s1Pos:] == s3[s3Pos:] {
			return TRUE97
		} else {
			return FALSE97
		}
	}

	result := byte(FALSE97)
	// 此时s1，s2，s3的都没有消费完
	for i := s1Pos; i < len(s1) && s3Pos+(i-s1Pos) < len(s3) && s1[i] == s3[s3Pos+(i-s1Pos)]; i++ {
		// 尝试1个字符到前缀相等的最长字符串
		s3Start := s3Pos + (i - s1Pos) + 1
		for j := s2Pos; j < len(s2) && s3Start+(j-s2Pos) < len(s3) && s2[j] == s3[s3Start+(j-s2Pos)]; j++ {
			res := isInterleaveDP(s1, s2, s3, i+1, j+1, s3Start+(j-s2Pos)+1, memo)
			if res == TRUE97 {
				result = TRUE97
				break
			}
		}
	}

	memo[s1Pos][s2Pos][s3Pos] = result
	return result
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	memo := make([][][]byte, len(s1)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([][]byte, len(s2)+1)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = make([]byte, len(s3)+1)
		}
	}

	if isInterleaveDP(s1, s2, s3, 0, 0, 0, memo) == TRUE97 {
		return true
	}
	// 重新初始化memo
	memo = make([][][]byte, len(s2)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([][]byte, len(s1)+1)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = make([]byte, len(s3)+1)
		}
	}
	if isInterleaveDP(s2, s1, s3, 0, 0, 0, memo) == TRUE97 {
		return true
	} else {
		return false
	}
}

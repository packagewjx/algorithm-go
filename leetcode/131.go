package leetcode

func partition(s string) [][]string {
	if s == "" {
		return [][]string{}
	}
	memo := make([][][]string, len(s)+1)
	memo[len(s)] = [][]string{{}}
	memo[len(s)-1] = [][]string{{s[len(s)-1:]}}

	var bt func(start int) [][]string
	bt = func(start int) [][]string {
		if memo[start] != nil {
			return memo[start]
		}

		result := make([][]string, 0, 10)

		// 确保i位置的回文串的最右边不超过s
		for i := start; i+(i-start) < len(s); i++ {
			// 以i为中心，长度为
			isHuiwen := true
			for j := 0; i-j >= start; j++ {
				if s[i-j] != s[i+j] {
					isHuiwen = false
					break
				}
			}
			if isHuiwen {
				// 说明s[start:i+(i-start)+1]是回文串
				huiwen := s[start : i+(i-start)+1]
				right := bt(i + (i - start) + 1)
				for j := 0; j < len(right); j++ {
					res := append([]string{huiwen}, right[j]...)
					result = append(result, res)
				}
			}

			// 以i与i+1为中心
			if i+(i-start)+1 < len(s) && s[i] == s[i+1] {
				isHuiwen = true
				// j从1开始的原因是，进来if语句已经判断了i与i+1位置的字符相等
				for j := 1; i-j >= start && i+1+j < len(s); j++ {
					if s[i-j] != s[i+1+j] {
						isHuiwen = false
						break
					}
				}

				if isHuiwen {
					huiwen := s[start : i+(i-start)+2]
					right := bt(i + (i - start) + 2)
					for j := 0; j < len(right); j++ {
						res := append([]string{huiwen}, right[j]...)
						result = append(result, res)
					}
				}
			}
		}

		memo[start] = result
		return result
	}

	return bt(0)
}

package leetcode

func balancedStringSplit(s string) int {
	numL := 0
	numR := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			numL++
		} else {
			numR++
		}

		if numL == numR {
			count++
			numL = 0
			numR = 0
		}
	}

	return count
}

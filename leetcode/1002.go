package leetcode

import "math"

func commonChars(A []string) []string {
	count := make([][]int, len(A))
	ret := make([]string, 0, 10)

	for i := 0; i < len(A); i++ {
		count[i] = make([]int, 26)

		for _, char := range A[i] {
			count[i][char-'a'] += 1
		}
	}

	for i := 0; i < 26; i++ {
		smallest := math.MaxInt64
		for j := 0; j < len(A); j++ {
			if count[j][i] < smallest {
				smallest = count[j][i]
			}
		}

		for smallest > 0 {
			ret = append(ret, string('a'+i))
			smallest--
		}
	}

	return ret
}

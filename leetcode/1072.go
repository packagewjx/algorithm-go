package leetcode

func maxEqualRowsAfterFlips(matrix [][]int) int {
	counts := make(map[string]int)

	for i := 0; i < len(matrix); i++ {
		key := ""
		if matrix[i][0] == 1 {
			strings := []string{"1", "0"}
			for j := 0; j < len(matrix[i]); j++ {
				key += strings[matrix[i][j]]
			}
		} else {
			strings := []string{"0", "1"}
			for j := 0; j < len(matrix[i]); j++ {
				key += strings[matrix[i][j]]
			}
		}
		counts[key] = counts[key] + 1
	}
	largest := 0
	for _, count := range counts {
		if count > largest {
			largest = count
		}
	}
	return largest
}

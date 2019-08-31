package leetcode

// 解法挺巧妙的
func largestOverlap(A [][]int, B [][]int) int {
	AOne := make([]int, 0)
	BOne := make([]int, 0)
	offset := 100
	N := len(A)
	for i := 0; i < N; i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 1 {
				AOne = append(AOne, i*offset+j)
			}
			if B[i][j] == 1 {
				BOne = append(BOne, i*offset+j)
			}
		}
	}

	count := make(map[int]int)
	for _, i := range AOne {
		for _, j := range BOne {
			count[i-j]++
		}
	}

	max := 0
	for _, v := range count {
		if v > max {
			max = v
		}
	}
	return max
}

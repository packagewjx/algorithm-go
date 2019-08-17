package leetcode

func minimumDeleteSum(s1 string, s2 string) int {
	// initialize
	deleteSum := make([][]int, len(s2)+1)
	for i := 0; i < len(deleteSum); i++ {
		deleteSum[i] = make([]int, len(s1)+1)
	}

	max := 0
	for i := 0; i < len(s1); i++ {
		max += int(s1[i])
	}
	for i := 0; i < len(s2); i++ {
		max += int(s2[i])
	}

	for i := 0; i <= len(s1); i++ {
		deleteSum[0][i] = max
	}
	for i := 0; i <= len(s2); i++ {
		deleteSum[i][0] = max
	}

	// do calculate
	for i := 0; i < len(s2); i++ {
		for j := 0; j < len(s1); j++ {
			if s2[i] == s1[j] {
				deleteSum[i+1][j+1] = deleteSum[i][j] - 2*int(s2[i])
			} else if deleteSum[i+1][j] < deleteSum[i][j+1] {
				deleteSum[i+1][j+1] = deleteSum[i+1][j]
			} else {
				deleteSum[i+1][j+1] = deleteSum[i][j+1]
			}
		}
	}

	return deleteSum[len(s2)][len(s1)]
}

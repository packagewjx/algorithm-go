package leetcode

// 最长子序列的另一个问法
func maxUncrossedLines(A []int, B []int) int {
	memo := make([][]int, len(A)+1)
	for i := 0; i < len(A)+1; i++ {
		memo[i] = make([]int, len(B)+1)
	}

	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				memo[i][j] = memo[i-1][j-1] + 1
			} else {
				if memo[i-1][j] > memo[i][j-1] {
					memo[i][j] = memo[i-1][j]
				} else {
					memo[i][j] = memo[i][j-1]
				}
			}
		}
	}
	return memo[len(A)][len(B)]
}

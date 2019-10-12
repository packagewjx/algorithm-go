package leetcode

func numBST(start, end int, memo []int) int {
	if end-start == -1 {
		return 1
	} else if memo[end-start] != -1 {
		return memo[end-start]
	}

	result := 0
	for i := start; i <= end; i++ {
		left := numBST(start, i-1, memo)
		right := numBST(i+1, end, memo)
		result += left * right
	}

	memo[end-start] = result
	return result
}

func numTrees(n int) int {
	memo := make([]int, n+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	memo[0] = 1
	memo[n] = 1

	return numBST(1, n, memo)
}

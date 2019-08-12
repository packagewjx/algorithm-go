package leetcode

func generateParenthesisRecursive(n int, memo *[][]string) []string {
	if (*memo)[n] != nil {
		return (*memo)[n]
	}

	set := make(map[string]bool, 0)
	for i := 1; i <= n-1; i++ {
		leftPerfect := generateParenthesisRecursive(i, memo)
		rightPerfect := generateParenthesisRecursive(n-i, memo)

		for j := 0; j < len(leftPerfect); j++ {
			for k := 0; k < len(rightPerfect); k++ {
				set[leftPerfect[j]+rightPerfect[k]] = true
			}
		}
	}

	// 除了上面的，还有最外面的括号配对的，中间n-1个括号的结果
	lastPartial := generateParenthesisRecursive(n-1, memo)
	for i := 0; i < len(lastPartial); i++ {
		set["("+lastPartial[i]+")"] = true
	}

	result := make([]string, 0, len(set))
	for str := range set {
		result = append(result, str)
	}

	// 记录结果
	(*memo)[n] = result
	// 返回
	return result
}

func generateParenthesis(n int) []string {
	memo := make([][]string, n+1)
	memo[1] = []string{"()"}
	memo[0] = []string{}
	return generateParenthesisRecursive(n, &memo)
}

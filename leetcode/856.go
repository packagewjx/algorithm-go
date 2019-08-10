package leetcode

func scoreOfParentheses(S string) int {
	levelScore := make([]int, 1, 10)
	curLevel := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			levelScore = append(levelScore, 0)
			curLevel++
		} else if S[i] == ')' {
			if S[i-1] == '(' {
				levelScore[curLevel-1] += 1
			} else {
				levelScore[curLevel-1] += 2 * levelScore[curLevel]
			}
			levelScore = levelScore[0 : len(levelScore)-1]
			curLevel--
		}
	}
	return levelScore[0]
}

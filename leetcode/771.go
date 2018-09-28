package leetcode

func numJewelsInStones(J string, S string) int {
	if len(J) == 0 || len(S) == 0 {
		return 0
	}

	result := 0

	//创建图来存储有哪些是宝石
	jewel := make(map[int32]bool)
	for _, c := range J {
		jewel[c] = true
	}

	for _, c := range S {
		if jewel[c] {
			result++
		}
	}

	return result
}

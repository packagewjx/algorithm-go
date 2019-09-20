package leetcode

func getDistanceDP(w1, w2 string, i, j int, memo [][]int) int {
	if i == len(w1) {
		return len(w2) - j
	} else if j == len(w2) {
		return len(w1) - i
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	// 实际计算编辑距离

	// 如果相等，则直接查看i+1和j+1即可
	if w1[i] == w2[j] {
		ret := getDistanceDP(w1, w2, i+1, j+1, memo)
		memo[i][j] = ret
		return ret
	}

	// 不相等的话，就要查看5种情况

	// 删除w1[i]或者在w2[j]前面加入字符w1[i]
	smallest := getDistanceDP(w1, w2, i+1, j, memo)
	// 删除w2[j]或者在w1[i]前面加入字符w2[j]
	temp := getDistanceDP(w1, w2, i, j+1, memo)
	if temp < smallest {
		smallest = temp
	}
	// 将w1[i]替换为w2[j]或者反过来
	temp = getDistanceDP(w1, w2, i+1, j+1, memo)
	if temp < smallest {
		smallest = temp
	}
	// 经过上述的操作之后，编辑距离加1
	smallest += 1
	memo[i][j] = smallest
	return smallest
}

func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(word2))
		for j := 0; j < len(word2); j++ {
			memo[i][j] = -1
		}
	}
	return getDistanceDP(word1, word2, 0, 0, memo)
}

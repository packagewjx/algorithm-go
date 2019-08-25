package leetcode

func removeStonesDFS(stones [][]int, cur int, visited []bool, next [][]int) int {
	if visited[cur] {
		return 0
	}
	count := 1
	visited[cur] = true
	for i := 0; i < len(next[cur]); i++ {
		count += removeStonesDFS(stones, next[cur][i], visited, next)
	}
	return count
}

func removeStones(stones [][]int) int {
	visited := make([]bool, len(stones))
	next := make([][]int, len(stones))
	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
				next[i] = append(next[i], j)
				next[j] = append(next[j], i)
			}
		}
	}

	count := 0
	for i := 0; i < len(stones); i++ {
		// 减1是因为不能加第i个，因为不能删除它
		if !visited[i] {
			count += removeStonesDFS(stones, i, visited, next) - 1
		}
	}
	return count
}

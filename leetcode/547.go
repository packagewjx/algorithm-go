package leetcode

func findCircleNum(M [][]int) int {
	visited := make([]bool, len(M))
	count := 0
	for i := 0; i < len(M); i++ {
		if visited[i] {
			continue
		}
		count++
		// 广度遍历
		queue := make([]int, 1)
		queue[0] = i
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if visited[cur] {
				continue
			}
			visited[cur] = true
			for j := 0; j < len(M); j++ {
				if M[cur][j] == 1 && j != cur {
					queue = append(queue, j)
				}
			}
		}
	}
	return count
}

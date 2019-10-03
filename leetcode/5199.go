package leetcode

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	if len(pairs) == 0 {
		return s
	}
	nextMap := make(map[int][]int)
	maxNode := 0
	for i := 0; i < len(pairs); i++ {
		nextMap[pairs[i][0]] = append(nextMap[pairs[i][0]], pairs[i][1])
		if maxNode < pairs[i][0] {
			maxNode = pairs[i][0]
		}
		nextMap[pairs[i][1]] = append(nextMap[pairs[i][1]], pairs[i][0])
		if maxNode < pairs[i][1] {
			maxNode = pairs[i][1]
		}
	}

	visited := make([]bool, maxNode+1)

	sets := make([][]int, 0, 10)
	for i := 0; i < len(visited); i++ {
		if visited[i] {
			continue
		}

		set := make([]int, 0, 10)
		// 广度遍历
		queue := make([]int, 0, 10)
		queue = append(queue, i)
		for len(queue) > 0 {
			n := queue[0]
			queue = queue[1:]
			if visited[n] {
				continue
			}
			visited[n] = true
			set = append(set, n)
			for j := 0; j < len(nextMap[n]); j++ {
				queue = append(queue, nextMap[n][j])
			}
		}
		sets = append(sets, set)
	}

	buf := []byte(s)
	for _, set := range sets {
		nums := make([]byte, len(set))
		for i := 0; i < len(set); i++ {
			nums[i] = s[set[i]]
		}
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		sort.Ints(set)
		for i := 0; i < len(set); i++ {
			buf[set[i]] = nums[i]
		}
	}

	return string(buf)
}

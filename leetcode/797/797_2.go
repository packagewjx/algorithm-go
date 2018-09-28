package _97

func allPathsSourceTarget_2(graph [][]int) [][]int {
	// 保存着所有从某个点能够去终点的路径
	pathMap := make(map[int][][]int)
	allPathRecursive(graph, 0, pathMap)
	return pathMap[0]
}

// 从终点出发，往前遍历，首先有一条路径，直接到终点的，返回这个路径，
// 然后查看前一个点，如果这个点能去刚刚返回的路径的第一个点，则
// 加入这个点到一条路径中，直到加到所有路径中，然后返回这些能去终点的路径
// 以此类推
func allPathRecursive(graph [][]int, pos int, pathMap map[int][][]int) {
	if pos == len(graph)-1 {
		pathMap[pos] = [][]int{{pos}}
		return
	}

	allPathRecursive(graph, pos+1, pathMap)

	newPathsToGoal := make([][]int, 0, 10)
	nexts := graph[pos]
	for _, next := range nexts {
		if len(pathMap[next]) > 0 {
			paths := pathMap[next]
			// 把当前点添加到所有路径中
			for _, path := range paths {
				newPath := make([]int, 1, len(path)+1)
				newPath[0] = pos
				newPath = append(newPath, path...)
				newPathsToGoal = append(newPathsToGoal, newPath)
			}
		}
	}
	pathMap[pos] = newPathsToGoal
}

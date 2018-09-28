package _97

var paths [][]int
var goal int

func allPathsSourceTarget_1(graph [][]int) [][]int {
	goal = len(graph) - 1
	paths = make([][]int, 0, 10)
	goRecursive(graph, 0, make([]int, 0, 10))
	return paths
}

func goRecursive(graph [][]int, pos int, path []int) {
	path = append(path, pos)
	if pos == goal {
		paths = append(paths, path)
	}

	next := graph[pos]
	if len(next) == 0 {
		return
	}

	// 走到任何的地方
	for _, val := range next {
		pathCopy := make([]int, len(path), len(path)+1)
		copy(pathCopy, path)
		goRecursive(graph, val, pathCopy)
	}
}

//TODO 可以进一步优化

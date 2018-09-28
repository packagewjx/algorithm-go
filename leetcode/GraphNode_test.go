package leetcode

import "testing"

func TestNewGraph(t *testing.T) {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	node := NewGraph(graph)
	print(node)
}

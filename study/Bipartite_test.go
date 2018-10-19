package study

import (
	"fmt"
	. "github.com/packagewjx/algorithm-go/datastructure"
	"testing"
)

func TestFindBipartite(t *testing.T) {
	graph := NewMatrixBasedGraphUsingList(map[int][]int{
		0: {5, 6},
		1: {5},
		2: {5, 7},
		3: {7, 8, 9},
		4: {8, 9},
		5: {0, 1, 2},
		6: {0},
		7: {2, 3},
		8: {3, 4},
		9: {3, 4}})

	u, v := findBipartiteVertices(graph)

	fmt.Println(u)
	fmt.Println(v)
}

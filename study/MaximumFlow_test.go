package study

import (
	"fmt"
	"github.com/packagewjx/algorithm-go/datastructure"
	"testing"
)

func TestMaximumFlow(t *testing.T) {
	matrix := [][]int{
		{0, 2, 0, 3, 0, 0},
		{0, 0, 5, 0, 3, 0},
		{0, 0, 0, 0, 0, 2},
		{0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 4},
		{0, 0, 0, 0, 0, 0}}

	graph := datastructure.BuildFromMatrix(matrix)

	fmt.Println(MaximumFlow(graph, graph[0], graph[5]))

}

package study

import (
	"fmt"
	"github.com/packagewjx/algorithm-go/datastructure"
	"testing"
)

func TestPrim(t *testing.T) {
	graph := &datastructure.MatrixBasedGraph{Matrix: [][]int{
		{0, 3, 0, 0, 6, 5},
		{3, 0, 1, 0, 0, 4},
		{0, 1, 0, 6, 0, 4},
		{0, 0, 6, 0, 8, 5},
		{6, 0, 0, 8, 0, 2},
		{5, 4, 4, 5, 2, 0}}}

	Prim(graph)
}

func TestKruscal(t *testing.T) {
	graph := &datastructure.MatrixBasedGraph{Matrix: [][]int{
		{0, 3, 0, 0, 6, 5},
		{3, 0, 1, 0, 0, 4},
		{0, 1, 0, 6, 0, 4},
		{0, 0, 6, 0, 8, 5},
		{6, 0, 0, 8, 0, 2},
		{5, 4, 4, 5, 2, 0}}}

	edgeSet := Kruscal(graph)
	fmt.Println(edgeSet)
}

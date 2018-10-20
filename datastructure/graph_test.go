package datastructure

import (
	"fmt"
	"testing"
)

func TestBuildFromMatrix(t *testing.T) {
	matrix := [][]int{
		{0, 2, 0, 3, 0, 0},
		{0, 0, 5, 0, 3, 0},
		{0, 0, 0, 0, 0, 2},
		{0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 4},
		{0, 0, 0, 0, 0, 0}}

	fromMatrix := BuildFromMatrix(matrix)
	fmt.Println(fromMatrix)
}

func TestEdgeSet_Add(t *testing.T) {
	edgeSet := NewEdgeSet()

	edgeSet.Add(0, 1)
	edgeSet.Add(0, 2)
	edgeSet.Add(0, 3)
	edgeSet.Add(0, 4)

	fmt.Println(edgeSet.AllEdges())
}

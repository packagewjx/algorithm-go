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

package study

import (
	"fmt"
	"github.com/packagewjx/algorithm-go/datastructure"
	"testing"
)

func TestBuildMap(t *testing.T) {
	testMap := map[int][]int{
		1:  {5, 6, 2},
		5:  {9, 10},
		9:  {5, 10},
		10: {9, 5},
		2:  {1, 6, 3},
		6:  {1, 2, 3},
		3:  {2, 6, 7, 4},
		7:  {3, 8},
		4:  {3, 8},
		8:  {4, 7}}

	vert := datastructure.BuildVerticesFromMap(testMap)

	points := ArticulationPoint(vert[1])
	fmt.Println(points)

	testMap = map[int][]int{
		1:  {2, 12, 3, 6},
		2:  {1, 3, 4, 7, 8, 13},
		3:  {1, 2},
		4:  {2, 5},
		5:  {4},
		6:  {1},
		7:  {2, 8, 11, 9},
		8:  {2, 7, 11},
		9:  {7},
		10: {12, 13},
		11: {7, 8},
		12: {1, 10, 13},
		13: {2, 10, 12}}
	vert = datastructure.BuildVerticesFromMap(testMap)
	points = ArticulationPoint(vert[1])
	fmt.Println(points)

}

func TestBuildMatrix(t *testing.T) {
	testMatrix := [][]bool{
		{false, true, false, false, true, true, false, false, false, false},
		{true, false, true, false, false, true, false, false, false, false},
		{false, true, false, true, false, false, true, false, false, false},
		{false, false, true, false, false, false, true, true, false, false},
		{true, false, false, false, false, false, false, false, true, true},
		{true, true, true, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, true, false, false},
		{false, false, true, true, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, true},
		{false, false, false, false, true, false, false, false, true, false}}

	matrix := datastructure.BuildVerticesFromMatrix(testMatrix)
	ArticulationPoint(matrix[1])
}

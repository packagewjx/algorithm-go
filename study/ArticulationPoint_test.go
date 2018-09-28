package study

import (
	"fmt"
	"testing"
)

func TestBuildMap(t *testing.T) {
	testMap := map[string][]string{
		"a": {"e", "f", "b"},
		"e": {"i", "j"},
		"i": {"e", "j"},
		"j": {"i", "e"},
		"b": {"a", "f", "c"},
		"f": {"a", "b", "c"},
		"c": {"b", "f", "g", "d"},
		"g": {"c", "h"},
		"d": {"c", "h"},
		"h": {"d", "g"}}

	vert := buildFromMap(testMap)

	points := articulationPoint(vert["a"])
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

	matrix := buildFromMatrix(testMatrix)
	fmt.Println(matrix)
}

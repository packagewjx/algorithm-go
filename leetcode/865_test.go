package leetcode

import "testing"

func Test865(t *testing.T) {
	node, _ := NewTree("[1,2,3,4,5,6,7]")
	deepest := subtreeWithAllDeepest(node)
	println(deepest)
}

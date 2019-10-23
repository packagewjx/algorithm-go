package leetcode

import "testing"

func Test124(t *testing.T) {
	node, _ := NewTree("[-10,9,20,null,null,15,7]")
	println(maxPathSum(node))
}

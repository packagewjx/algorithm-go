package leetcode

import "testing"

func Test99(t *testing.T) {
	node, _ := NewTree("[1,3,null,null,2]")
	recoverTree(node)
}

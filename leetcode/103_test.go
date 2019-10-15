package leetcode

import "testing"

func Test103(t *testing.T) {
	node, _ := NewTree("[1,2,3,4,5,6,7,8,9]")
	zigzagLevelOrder(node)
}

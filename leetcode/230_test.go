package leetcode

import "testing"

func Test230(t *testing.T) {
	node, _ := NewTree("[5,3,6,2,4,null,null,1]")
	println(kthSmallest(node, 5))
}

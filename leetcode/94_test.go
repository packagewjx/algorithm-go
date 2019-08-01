package leetcode

import "testing"

func Test94(t *testing.T) {
	array, _ := newTreeArray("[1,null,2,4,3,null,5]")
	root, _ := construct(array)
	traversal := inorderTraversal(root)
	println(traversal)
}

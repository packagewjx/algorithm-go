package leetcode

import "testing"

func Test5264(t *testing.T) {
	//node, _ := NewTree("[-1,null,-1]")
	constructor := Constructor(&TreeNode{Val: 0})
	constructor.Find(2)
	constructor.Find(0)
}

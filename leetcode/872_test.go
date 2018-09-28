package leetcode

import (
	"testing"
)

func Test872(t *testing.T) {
	treeNode, _ := NewTree("[3,5,1,6,2,9,8,null,null,7,4]")
	leafSimilar(treeNode, treeNode)
}

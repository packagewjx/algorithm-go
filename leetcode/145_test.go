package leetcode

import (
	"fmt"
	"testing"
)

func Test145(t *testing.T) {
	node, _ := NewTree("[1,null,2,3]")
	fmt.Println(postorderTraversal(node))
}

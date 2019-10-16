package leetcode

import (
	"fmt"
	"testing"
)

func Test113(t *testing.T) {
	node, e := NewTree("[-2,null,-3]")
	if e != nil {
		panic(e)
	}
	fmt.Println(pathSum(node, -5))
}

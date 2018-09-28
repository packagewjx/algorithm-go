package leetcode

import (
	"fmt"
	"testing"
)

func Test876(t *testing.T) {
	list := NewList([]int{1, 2, 3, 4, 5, 6})
	node := middleNode(list)
	fmt.Println(node)
}

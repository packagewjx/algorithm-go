//+build 86

package leetcode

import (
	"fmt"
	"testing"
)

func Test86(t *testing.T) {
	node := partition(NewList([]int{777, 7, 1}), 6)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

}

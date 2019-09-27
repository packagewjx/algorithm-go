package leetcode

import (
	"fmt"
	"testing"
)

func Test25(t *testing.T) {
	list := reverseKGroup(NewList([]int{1, 2, 3, 4, 5, 6}), 2)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

package leetcode

import "testing"

func Test23(t *testing.T) {
	l1 := NewList([]int{1, 4, 5})
	l2 := NewList([]int{1, 3, 4})
	l3 := NewList([]int{2, 6})
	mergeKLists([]*ListNode{l1, l2, l3})
}

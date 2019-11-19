package leetcode

import "testing"

func Test142(t *testing.T) {
	n1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	n2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	n1.Next = n2
	n2.Next = n1
	println(detectCycle(n1).Val)
}

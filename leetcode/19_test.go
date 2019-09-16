package leetcode

import "testing"

func Test19(t *testing.T) {
	removeNthFromEnd(&ListNode{
		Val:  1,
		Next: nil,
	}, 2)
}

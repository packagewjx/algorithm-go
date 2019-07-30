package leetcode

import "testing"

func Test515(t *testing.T) {
	values := largestValues(&TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	})
	print(values)
}

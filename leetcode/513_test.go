package leetcode

import "testing"

func Test513(t *testing.T) {
	array, _ := newTreeArray("[1,2,3,4,null,5,6,null,null,7]")
	root, _ := construct(array)

	findBottomLeftValue(root)
}

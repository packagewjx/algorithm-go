//+build 919

package leetcode

import "testing"

func Test919(t *testing.T) {
	node, _ := NewTree("[1,2,3,4,5,6]")
	inserter := Constructor(node)
	println(inserter.Insert(7))
	println(inserter.Insert(8))
	println(inserter.Insert(9))
}

package leetcode

import "testing"

func Test946(t *testing.T) {
	test := []int{1, 2, 5, 8, 7, 4, 6, 9, 3, 0}
	popped := []int{2, 7, 8, 4, 0, 5, 9, 1, 3, 6}
	println(validateStackSequences(test, popped))
}

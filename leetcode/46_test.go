package leetcode

import "testing"

func TestReverseSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	reverseSlice(slice)
	println(slice)
}

func Test46(t *testing.T) {
	permute([]int{1, 2, 3, 4})
}

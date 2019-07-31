package leetcode

import "testing"

func Test260(t *testing.T) {
	test := []int{1, 2, 1, 3, 2, 5}
	number := singleNumber(test)
	println(number[0])
	println(number[1])
}

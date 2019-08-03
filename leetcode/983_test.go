package leetcode

import "testing"

func Test983(t *testing.T) {
	days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	costs := []int{2, 7, 15}
	println(mincostTickets(days, costs))
}

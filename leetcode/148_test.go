package leetcode

import (
	"fmt"
	"testing"
)

func Test148(t *testing.T) {
	list := sortList(NewList([]int{1, 2, 0}))
	fmt.Println(list)
}

package leetcode

import (
	"fmt"
	"testing"
)

func Test68(t *testing.T) {
	justify := fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16)
	for i := 0; i < len(justify); i++ {
		fmt.Println(justify[i])
	}
}

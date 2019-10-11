package leetcode

import (
	"fmt"
	"testing"
)

func Test89(t *testing.T) {
	code := grayCode(4)
	for _, val := range code {
		fmt.Printf(fmt.Sprintf("%%0%db\n", 4), val)
	}
}

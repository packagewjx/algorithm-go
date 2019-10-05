//+build 51

package leetcode

import (
	"fmt"
	"testing"
)

func Test51(t *testing.T) {
	queens := solveNQueens(4)
	for i := 0; i < len(queens); i++ {
		fmt.Println()
		for j := 0; j < len(queens[i]); j++ {
			fmt.Println(queens[i][j])
		}
		fmt.Println()
	}
}

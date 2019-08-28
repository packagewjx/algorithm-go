package leetcode

import (
	"fmt"
	"testing"
)

func Test529(t *testing.T) {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'M'},
		{'E', 'E', 'M', 'E', 'E', 'E', 'E', 'E'},
		{'M', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'M', 'E', 'E', 'E', 'E'}}

	updateBoard(board, []int{0, 0})

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

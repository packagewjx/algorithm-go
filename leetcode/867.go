package leetcode

func transpose(A [][]int) [][]int {
	B := make([][]int, len(A[0]))
	for i := 0; i < len(A[0]); i++ {
		B[i] = make([]int, len(A))
	}

	for i, row := range A {
		for j, val := range row {
			B[j][i] = val
		}
	}

	return B
}

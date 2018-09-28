package leetcode

func flipAndInvertImage(A [][]int) [][]int {
	reverse := make([][]int, len(A))

	for i := 0; i < len(A); i++ {
		reverse[i] = make([]int, len(A[i]))
		array := reverse[i]

		for j := 0; j < len(array); j++ {
			if A[i][len(A[i])-j-1] == 0 {
				array[j] = 1
			} else {
				array[j] = 0
			}
		}
	}

	return reverse
}

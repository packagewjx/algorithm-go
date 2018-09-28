package leetcode

func peakIndexInMountainArray(A []int) int {
	var i int
	for i = 0; A[i] < A[i+1]; i++ {
	}
	return i
}

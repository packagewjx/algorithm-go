package leetcode

var temp = []int{0, 1, 2, 6}

func clumsy(N int) int {
	if N < 4 {
		return temp[N]
	}
	result := N*(N-1)/(N-2) + (N - 3)
	for N = N - 4; N >= 4; N -= 4 {
		result -= N * (N - 1) / (N - 2)
		result += N - 3
	}
	result -= temp[N]
	return result
}

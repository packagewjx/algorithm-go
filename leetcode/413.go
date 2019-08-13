package leetcode

func numberOfArithmeticSlices(A []int) int {
	result := 0
	for i := 0; i < len(A)-2; i++ {
		// 寻找第一个等差数列
		for ; i < len(A)-2 && A[i]-A[i+1] != A[i+1]-A[i+2]; i++ {
		}
		if i == len(A)-2 {
			break
		}

		j := i + 1
		// 现在，A[i], A[i+1], A[i+2]构成等差数列，现在就要找出等差数列的结束
		for ; j < len(A)-2 && A[j]-A[j+1] == A[j+1]-A[j+2]; j++ {

		}
		// 那么，这里A[i], A[i+1], ..., A[j], A[j+1]构成等差数列

		// 等差数列加法公式
		result += (1 + (j - i)) * (j - i) / 2
		// 继续从j+1处寻找等差数列，因为下一步会+1
		i = j
	}

	return result
}

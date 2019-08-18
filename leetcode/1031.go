package leetcode

func maxSumTwoNoOverlap(A []int, L int, M int) int {
	// 前缀和
	for i := 1; i < len(A); i++ {
		A[i] += A[i-1]
	}
	res := A[L+M-1]
	// 初始化，认为从第一个开始的子数组是最大的
	Lmax := A[L-1]
	Mmax := A[M-1]
	max := func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		} else {
			return i2
		}
	}

	for i := L + M; i < len(A); i++ {
		// 每一轮，LMax都更新为从i-L-M开始的L个数字的最大，也就是[0,i-L-M]中最大的L长度的子数组
		Lmax = max(Lmax, A[i-M]-A[i-L-M])
		// 同上
		Mmax = max(Mmax, A[i-L]-A[i-L-M])

		// 然后，这个最长的L的子数组加上从i-M开始的M长度的子数组，这样L和M子数组是不重合的，并且取得了L的最大值。
		// 没有看M是否最大，但是若是，则res就成了最大的了
		res = max(res, max(Lmax+A[i]-A[i-M], Mmax+A[i]-A[i-L]))
	}
	return res
}

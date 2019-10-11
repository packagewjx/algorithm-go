package leetcode

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	buf := make([][]int, 2)
	buf[1] = make([]int, 2, 1<<uint(n))
	buf[1][0] = 0
	buf[1][1] = 1
	buf[0] = make([]int, 0, 1<<uint(n))

	for i := 2; i <= n; i++ {
		last := &buf[1-i&1]
		this := &buf[i&1]
		*this = (*this)[:0]

		// last中的数都是i-1位的，共2^(j-1)个
		// 正序遍历last，在前面加上0，复制就行
		for j := 0; j < len(*last); j++ {
			*this = append(*this, (*last)[j])
		}

		// 反序遍历last，在前面加上1
		mask := 1 << uint(i-1)
		for j := len(*last) - 1; j >= 0; j-- {
			*this = append(*this, (*last)[j]|mask)
		}
	}

	return buf[n&1]
}

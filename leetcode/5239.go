package leetcode

func circularPermutation(n int, start int) []int {
	var buf [2][]int
	buf[0] = []int{0}
	startPos := 0
	for i := 1; i <= n; i++ {
		last := &buf[i&1^1]
		this := &buf[i&1]
		*this = (*this)[:0]
		for j := 0; j < len(*last); j++ {
			*this = append(*this, (*last)[j])
			if n == i && (*this)[len(*this)-1] == start {
				startPos = len(*this) - 1
			}
		}
		mask := 1 << uint(i-1)
		for j := len(*last) - 1; j >= 0; j-- {
			*this = append(*this, (*last)[j]|mask)
			if n == i && (*this)[len(*this)-1] == start {
				startPos = len(*this) - 1
			}
		}
	}

	this := &buf[n&1]
	res := append((*this)[startPos:], (*this)[0:startPos]...)
	return res
}

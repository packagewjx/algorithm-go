package leetcode

func constructArray(n int, k int) []int {
	tail := make([]int, k+1)
	for i := 0; i < len(tail); i += 2 {
		tail[i] = i/2 + 1
	}
	for i := 1; i < len(tail); i += 2 {
		tail[i] = k + 1 - (i / 2)
	}

	add := n - k - 1
	for i := 0; i < len(tail); i++ {
		tail[i] += add
	}
	ret := make([]int, n)
	for i := 0; i < tail[0]; i++ {
		ret[i] = i + 1
	}
	copy(ret[tail[0]-1:], tail)
	return ret
}

//+build 5198

package leetcode

func nthUglyNumber(n int, a int, b int, c int) int {
	if n == 1 {
		return 1
	}

	cur := 0
	an := a
	bn := b
	cn := c
	num := 1
	for ; cur < n-1; num++ {
		found := false
		if num == an {
			found = true
			an += a
		}
		if num == bn {
			found = true
			bn += b
		}
		if num == cn {
			found = true
			cn += c
		}
		if found {
			cur++
		}
	}

	return num - 1
}

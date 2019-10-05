package leetcode

func min3Num(n1, n2, n3 int) int {
	if n1 < n2 {
		if n1 < n3 {
			return n1
		} else /*n1 >= n3 */ {
			return n3
		}
	} else /*n1 >= n2*/ {
		if n2 < n3 {
			return n2
		} else /*n2 >= n3*/ {
			return n3
		}
	}
}

// 使用容斥原理加二分查找寻找第n个丑数
func nthUglyNumber(n int, a int, b int, c int) int {
	l := 1
	r := min3Num(n*a, n*b, n*c)
	ab := a * b
	ac := a * c
	bc := b * c
	abc := a * b * c

	for l <= r {
		m := (l + r) / 2

		// 容斥原理计算m是第几个
		num := m/a + m/b + m/c - m/ab - m/ac - m/bc + m/abc
		if num == n {
			return m
		} else if num < n {
			l = m + 1
		} else /* num > n*/ {
			r = m - 1
		}
	}
	panic("impossible")
}

func gcd(a, b int) int {
	if a < b {
		temp := a
		a = b
		b = temp
	}

	c := a % b
	for c > 0 {
		a = b
		b = c
		c = a % b
	}
	return b
}

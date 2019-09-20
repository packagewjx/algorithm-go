package leetcode

import "sort"

func myPow(x float64, n int) float64 {
	// 处理特殊情况
	if n == 0 || x == 1 {
		return 1
	} else if n == 1 {
		return x
	} else if n == -1 {
		return 1 / x
	} else if x == -1 {
		t := n & 1
		if t == 1 {
			return -1
		} else {
			return 1
		}
	}

	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	midRes := make([]float64, 1, 10)
	powOf2 := make([]int, 1, 10)
	midRes[0] = x
	powOf2[0] = 1
	mi := 1
	num := x
	for mi < n {
		num = num * num
		mi <<= 1
		powOf2 = append(powOf2, mi)
		midRes = append(midRes, num)
	}
	// midRes长度一定大于等于2
	res := midRes[len(midRes)-2]
	//加入剩下的幂
	mi = n - powOf2[len(powOf2)-2]
	for mi > 0 {
		pos := sort.SearchInts(powOf2, mi)
		// pos一定大于等于1
		if mi == powOf2[pos] {
			res *= midRes[pos]
			break
		} else {
			pos--
			mi -= powOf2[pos]
			res *= midRes[pos]
		}
	}

	if negative {
		return 1 / res
	} else {
		return res
	}
}

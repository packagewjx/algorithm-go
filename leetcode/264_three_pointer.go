//+build 264tp

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

func nthUglyNumber(n int) int {
	uglyNums := make([]int, 1, n+1)
	uglyNums[0] = 1
	i1, i2, i3 := 0, 0, 0
	for i := 1; i < n; i++ {
		min := min3Num(uglyNums[i1]*2, uglyNums[i2]*3, uglyNums[i3]*5)
		uglyNums = append(uglyNums, min)
		if min == uglyNums[i1]*2 {
			i1++
		}
		if min == uglyNums[i2]*3 {
			i2++
		}
		if min == uglyNums[i3]*5 {
			i3++
		}
	}
	return uglyNums[len(uglyNums)-1]
}

package contest

func fractionRecur(cont []int) []int {
	if len(cont) == 1 {
		return []int{1, cont[0]}
	}

	recur := fractionRecur(cont[1:])
	return []int{recur[1], recur[1]*cont[0] + recur[0]}
}

func maxGongyue(a, b int) int {
	var max, min int
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	c := max % min
	if c == 0 {
		return min
	}
	return maxGongyue(c, min)
}

func fraction(cont []int) []int {
	result := fractionRecur(cont)

	// 求两个的最大公约数
	gongyue := maxGongyue(result[0], result[1])
	// 多转了一次，因此掉转
	return []int{result[1] / gongyue, result[0] / gongyue}
}

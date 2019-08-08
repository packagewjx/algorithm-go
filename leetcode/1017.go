package leetcode

func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}

	// 判断是否是负奇数
	mask := uint(0x8000000000000001)
	str := []string{"0", "1"}
	result := ""
	// 使用特殊短除法，永远保证余数是1或0，而不出现-1
	for N != 0 {
		if uint(N)&mask == mask {
			// 是负奇数
			result = "1" + result
			N = N/-2 + 1
		} else {
			// 正常处理
			result = str[N%-2] + result
			N = N / -2
		}
	}
	return result
}

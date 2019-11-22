package leetcode

import "math"

func myAtoi(str string) int {
	const max = math.MaxInt32 / 10
	negative := false
	cur := 0
	for ; cur < len(str) && str[cur] == ' '; cur++ {
	}
	// 若全是空格，则返回
	if cur == len(str) {
		return 0
	}
	if str[cur] == '-' {
		negative = true
		cur++
	} else if str[cur] == '+' {
		cur++
	}
	if cur == len(str) || !(str[cur] <= '9' && str[cur] >= '0') {
		// 到结尾了或者非数字
		return 0
	}
	// 转换第一个数字
	num := int32(str[cur] - '0')
	cur++
	yuejie := false
	// 转换数字
	for ; cur < len(str) && str[cur] >= '0' && str[cur] <= '9'; cur++ {
		// 判断越界
		if max < num || (max == num && str[cur]-'0' > 7) {
			yuejie = true
		}
		num = num*10 + int32(str[cur]-'0')
	}
	if yuejie {
		if negative {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	} else {
		if negative {
			return -int(num)
		} else {
			return int(num)
		}
	}
}

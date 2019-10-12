package leetcode

import (
	"strconv"
	"strings"
)

func restoreIpAddressesBT(s string, cur []int, pos int, result *[]string) {
	if len(cur) == 4 {
		if pos == len(s) {
			// 转换为字符串
			builder := strings.Builder{}
			builder.WriteString(strconv.Itoa(cur[0]))
			builder.WriteString(".")
			builder.WriteString(strconv.Itoa(cur[1]))
			builder.WriteString(".")
			builder.WriteString(strconv.Itoa(cur[2]))
			builder.WriteString(".")
			builder.WriteString(strconv.Itoa(cur[3]))
			*result = append(*result, builder.String())
		}
		return
	} else if pos == len(s) {
		// 这里是没有凑够4个数字的
		return
	} else if len(s)-pos > 3*(4-len(cur)) {
		// 剩下的数字，就算3个3个转换，都花费不完，则可以返回了
		return
	}

	// 尝试1数字
	n := int(s[pos] & 15)
	try := make([]int, len(cur))
	copy(try, cur)
	try = append(try, n)
	restoreIpAddressesBT(s, try, pos+1, result)
	// 如果是0，则只允许使用1个数字
	if n == 0 {
		return
	}
	// 2数字
	if pos+1 < len(s) {
		n = 10*n + int(s[pos+1]&15)
		try := make([]int, len(cur))
		copy(try, cur)
		try = append(try, n)
		restoreIpAddressesBT(s, try, pos+2, result)
	}
	// 3数字
	if pos+2 < len(s) {
		n = 10*n + int(s[pos+2]&15)
		if n <= 255 {
			try := make([]int, len(cur))
			copy(try, cur)
			try = append(try, n)
			restoreIpAddressesBT(s, try, pos+3, result)
		}
	}
}

func restoreIpAddresses(s string) []string {
	result := make([]string, 0, 10)
	restoreIpAddressesBT(s, []int{}, 0, &result)
	return result
}

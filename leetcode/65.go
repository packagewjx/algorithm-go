package leetcode

import "strings"

func isNumber(s string) bool {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return false
	}
	pos := 0
	// 符号位判断
	if s[pos] == '-' || s[pos] == '+' {
		pos++
	}
	oldPos := pos

	for pos < len(s) && s[pos] <= '9' && s[pos] >= '0' {
		pos++
	}
	if pos < len(s) && s[pos] == '.' {
		if pos == oldPos {
			// 说明前面没有数字，此时不允许单独一个点的情况
			pos++
			if pos == len(s) || !(s[pos] <= '9' && s[pos] >= '0') {
				return false
			}
		}
		pos++
		for pos < len(s) && s[pos] <= '9' && s[pos] >= '0' {
			pos++
		}
	}
	if oldPos == pos {
		// 位置没变动，说明没有数字，不能再继续了
		return false
	}
	if pos < len(s) && s[pos] == 'e' {
		pos++
		if pos < len(s) && (s[pos] == '-' || s[pos] == '+') {
			pos++
		}
		if pos == len(s) || !(s[pos] <= '9' && s[pos] >= '0') {
			return false
		}
		pos++
		for pos < len(s) && s[pos] <= '9' && s[pos] >= '0' {
			pos++
		}
	}

	return pos == len(s)
}

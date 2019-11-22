package leetcode

import "strings"

func reverseWords(s string) string {
	builder := strings.Builder{}

	wordEnd := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if wordEnd != -1 {
				builder.WriteString(s[i+1 : wordEnd])
				builder.WriteByte(' ')
				wordEnd = -1
			}
		} else if wordEnd == -1 {
			wordEnd = i + 1
		}
	}
	if wordEnd != -1 {
		builder.WriteString(s[0:wordEnd])
	}

	str := builder.String()
	// 多了个空格的话，要去掉
	if len(str) > 0 && str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}
	return str
}

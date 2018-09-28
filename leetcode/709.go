package leetcode

import (
	"bytes"
)

func toLowerCase(str string) string {
	diff := 'a' - 'A'
	Bytes := make([]byte, len(str))

	for i, char := range str {
		// 注意，这个转换只能是在字母范围！
		if char <= 'Z' && char >= 'A' {
			Bytes[i] = byte(char + diff)
		} else {
			Bytes[i] = byte(char)
		}
	}
	buf := bytes.NewBuffer(Bytes)
	return buf.String()
}

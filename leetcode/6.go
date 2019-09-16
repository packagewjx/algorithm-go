package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	totalLength := numRows + (len(s)/(numRows*2-2))*(numRows*2-2)
	result := make([]byte, 0, totalLength)

	// 第一行
	factor := 2*numRows - 2
	pos := 0
	index := 0
	for pos < len(s) {
		result = append(result, s[pos])
		pos += factor
		index++
	}
	// 第二到第numRows-1行 LDREOEIIECIHNTSG
	for r := 1; r < numRows-1; r++ {
		choice := false
		pos = r
		c1 := r
		c2 := 2*numRows - 2 - r
		for pos < len(s) {
			if !choice {
				// 加上竖的
				result = append(result, s[c1])
				c1 += factor
				pos = c2
			} else {
				// 加上斜的
				result = append(result, s[c2])
				c2 += factor
				pos = c1
			}
			choice = !choice
		}
	}
	// 最后一行
	pos = numRows - 1
	for pos < len(s) {
		result = append(result, s[pos])
		pos += factor
	}
	return string(result)
}

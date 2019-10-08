package leetcode

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0, 10)
	lineWords := []string{words[0]}
	wordLen := len(words[0])
	for i := 1; i < len(words); i++ {
		// 判断单词中间都有一个空格的时候，是否等于maxWidth
		if wordLen+len(words[i])+len(lineWords) <= maxWidth {
			lineWords = append(lineWords, words[i])
			wordLen += len(words[i])
		} else {
			// 这一行满了
			builder := strings.Builder{}
			if len(lineWords) == 1 {
				builder.WriteString(lineWords[0])
				for builder.Len() < maxWidth {
					builder.WriteString(" ")
				}
			} else {
				space := maxWidth - wordLen
				eachSpace := space / (len(lineWords) - 1)
				spaces := make([]byte, eachSpace)
				for k := 0; k < eachSpace; k++ {
					spaces[k] = ' '
				}
				addSpace := space % (len(lineWords) - 1)
				builder.WriteString(lineWords[0])
				for j := 1; j < len(lineWords); j++ {
					if addSpace > 0 {
						builder.WriteString(" ")
						addSpace--
					}
					builder.Write(spaces)
					builder.WriteString(lineWords[j])
				}
			}

			res = append(res, builder.String())

			// 重新初始化
			lineWords = []string{words[i]}
			wordLen = len(words[i])
		}
	}

	// 处理最后一行
	builder := strings.Builder{}
	builder.WriteString(lineWords[0])
	for i := 1; i < len(lineWords); i++ {
		builder.WriteString(" ")
		builder.WriteString(lineWords[i])
	}
	for builder.Len() < maxWidth {
		builder.WriteString(" ")
	}

	res = append(res, builder.String())
	return res
}

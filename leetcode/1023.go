package leetcode

func camelMatch(queries []string, pattern string) []bool {
	words := make([]string, 0)
	for i := 0; i < len(pattern); i++ {
		if pattern[i] >= 'A' && pattern[i] <= 'Z' {
			end := i + 1
			for ; end < len(pattern); end++ {
				if pattern[end] >= 'A' && pattern[end] <= 'Z' {
					break
				}
			}
			words = append(words, pattern[i:end])
			// 下一步++，因此要减1
			i = end - 1
		}
	}

	result := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		checked := 0
		isFalse := false
		for j := 0; j < len(words); j++ {
			// 将checked移动到大写字母处
			for ; checked < len(query) && query[checked] >= 'a' && query[checked] <= 'z'; checked++ {
			}
			if checked == len(query) {
				// 还没检查完就已经结束了，因此不匹配
				isFalse = true
				break
			}
			word := words[j]

			// 检查单词是否出现
			for k := 0; k < len(word); k++ {
				for ; checked < len(query); checked++ {
					if query[checked] == word[k] {
						// 检查通过，检查下一个
						checked++
						break
					} else if query[checked] >= 'A' && query[checked] <= 'Z' {
						// 这里进入了下一个单词，意味着是错误的
						isFalse = true
						break
					}
				}
			}
			if isFalse {
				break
			}
		}
		if !isFalse {
			// 查看接下来还有没有单词，也就是查看是否还有大写字母
			for ; checked < len(query) && query[checked] >= 'a' && query[checked] <= 'z'; checked++ {
			}
			if checked == len(query) {
				result[i] = true
			}
		}
	}

	return result
}

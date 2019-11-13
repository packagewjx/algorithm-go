package leetcode

import "strings"

func wordBreak(s string, wordDict []string) []string {
	if len(wordDict) == 0 {
		return []string{}
	}

	wordMap := make(map[string]bool)
	longest := 0
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = true
		if len(wordDict[i]) > longest {
			longest = len(wordDict[i])
		}
	}

	// 首先判断是否可以拆分
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j <= len(s) && j-i <= longest; j++ {
			if wordMap[s[i:j]] {
				if dp[j] {
					dp[i] = true
					break
				}
			}
		}
	}

	if !dp[0] {
		return []string{}
	}

	type context struct {
		ends []int
	}
	queue := make([]*context, 0, 100)
	queue = append(queue, &context{
		ends: []int{0},
	})

	result := make([]string, 0, 10)
	for len(queue) > 0 {
		ctx := queue[0]
		queue = queue[1:]

		end := ctx.ends[len(ctx.ends)-1]
		for i := end + 1; i-end <= longest && i <= len(s); i++ {
			if wordMap[s[end:i]] {
				if i == len(s) {
					// 最后一个单词。开始取得结果
					builder := strings.Builder{}
					for j := 1; j < len(ctx.ends); j++ {
						builder.WriteString(s[ctx.ends[j-1]:ctx.ends[j]])
						builder.WriteByte(' ')
					}
					builder.WriteString(s[ctx.ends[len(ctx.ends)-1]:])
					result = append(result, builder.String())
				} else {
					newEnds := make([]int, len(ctx.ends)+1)
					copy(newEnds, ctx.ends)
					newEnds[len(newEnds)-1] = i
					queue = append(queue, &context{ends: newEnds})
				}
			}
		}
	}
	return result
}

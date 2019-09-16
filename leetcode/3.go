package leetcode

/*
记录每一个字符的最多两个位置，到pos数组中。最理想的情况，记当前位置i的字符为c，若记录了c两次出现的位置pos[c][0]与pos[c][1]，则最长的不重复字符串是
[pos[c][0]+1,i]的子字符串。但是也要考虑其他字符的重复情况，因此使用一个变量startPos，记录着不重复的子字符串的最后的开始位置。途中不断让pos[c][0]
与startPos进行比较，并更新startPos的值，以此获得结果.
*/
func lengthOfLongestSubstring(s string) int {
	pos := make([][]int, 128)

	largest := 0
	// 记录着不重复子串的开始位置
	startPos := 0

	for i := 0; i < len(s); i++ {
		siLen := len(pos[s[i]])
		if siLen == 0 {
			pos[s[i]] = append(pos[s[i]], i)
		} else if siLen == 1 {
			// 若前面有一个，则可以取到startPos到i的子串
			l := i - startPos
			if l > largest {
				largest = l
			}
			pos[s[i]] = append(pos[s[i]], i)
			// 需要更新为更后面的那一个，防止重复
			if pos[s[i]][0]+1 > startPos {
				startPos = pos[s[i]][0] + 1
			}
		} else /* siLen == 2 */ {
			// 这里是只允许有2个的情况
			l := 0
			if startPos > pos[s[i]][0]+1 {
				l = i - startPos
			} else {
				l = i - pos[s[i]][0] - 1
			}
			if l > largest {
				largest = l
			}
			pos[s[i]] = append(pos[s[i]], i)
			pos[s[i]] = pos[s[i]][1:]
			startPos = pos[s[i]][0] + 1
		}
	}
	if len(s)-startPos > largest {
		largest = len(s) - startPos
	}
	return largest
}

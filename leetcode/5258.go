package leetcode

func maxScoreWords(words []string, letters []byte, score []int) int {
	letterMap := make([]int, 128)
	for i := 0; i < len(letters); i++ {
		letterMap[letters[i]]++
	}
	wordLetterMap := make([][]int, len(words))
	wordScore := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		m := make([]int, 128)
		s := 0
		for j := 0; j < len(words[i]); j++ {
			m[words[i][j]]++
			s += score[words[i][j]-'a']
		}
		wordLetterMap[i] = m
		wordScore[i] = s
	}

	canUse := func(wordIdx int, letterMap []int) bool {
		m := wordLetterMap[wordIdx]
		for i := 'a'; i <= 'z'; i++ {
			if m[i] > letterMap[i] {
				return false
			}
		}
		return true
	}

	max := 0
	var bt func(curScore int, pos int)
	bt = func(curScore int, pos int) {
		if pos == len(words) {
			if curScore > max {
				max = curScore
			}
			return
		}
		// 用的情况
		if canUse(pos, letterMap) {
			for j := 'a'; j <= 'z'; j++ {
				letterMap[j] -= wordLetterMap[pos][j]
			}
			curScore += wordScore[pos]
			bt(curScore, pos+1)
			// 复原
			for j := 'a'; j <= 'z'; j++ {
				letterMap[j] += wordLetterMap[pos][j]
			}
			curScore -= wordScore[pos]
		}
		// 不用的情况
		bt(curScore, pos+1)
	}
	bt(0, 0)

	return max
}

package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string][]string)
	foundEnd := false
	for i := 0; i < len(wordList); i++ {
		if endWord == wordList[i] {
			foundEnd = true
		}
		buf := []byte(wordList[i])
		for j := 0; j < len(buf); j++ {
			ori := buf[j]
			buf[j] = '*'
			key := string(buf)
			wordMap[key] = append(wordMap[key], wordList[i])
			buf[j] = ori
		}
	}
	if !foundEnd {
		return 0
	}

	type bfsContext struct {
		word  string
		count int
	}

	// 双向广度优先遍历
	forwardVisited := make(map[string]int)
	forwardVisited[beginWord] = 1
	backwardVisited := make(map[string]int)
	backwardVisited[endWord] = 1
	forwardQueue := make([]*bfsContext, 0, 100)
	backwardQueue := make([]*bfsContext, 0, 100)
	forwardQueue = append(forwardQueue, &bfsContext{
		word:  beginWord,
		count: 1,
	})
	backwardQueue = append(backwardQueue, &bfsContext{
		word:  endWord,
		count: 1,
	})

	visitNode := func(queue *[]*bfsContext, visited map[string]int, otherVisited map[string]int) int {
		c := (*queue)[0]
		*queue = (*queue)[1:]

		//visited[c.word] = c.count

		buf := []byte(c.word)
		for i := 0; i < len(c.word); i++ {
			ori := buf[i]
			buf[i] = '*'
			key := string(buf)
			for _, word := range wordMap[key] {
				if otherVisited[word] != 0 {
					// 前后向都找到了
					return c.count + otherVisited[word]
				}

				if visited[word] == 0 {
					visited[word] = c.count + 1
					*queue = append(*queue, &bfsContext{
						word:  word,
						count: c.count + 1,
					})
				}
			}
			buf[i] = ori
		}
		return -1
	}

	for len(forwardQueue) > 0 && len(backwardQueue) > 0 {
		node := visitNode(&forwardQueue, forwardVisited, backwardVisited)
		if node != -1 {
			return node
		}

		node = visitNode(&backwardQueue, backwardVisited, forwardVisited)
		if node != -1 {
			return node
		}
	}

	return 0
}

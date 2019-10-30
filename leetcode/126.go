package leetcode

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
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
		return [][]string{}
	}

	type bfsContext struct {
		word  string
		count int
	}

	type entry struct {
		last  []string
		count int
	}

	visited := make(map[string]*entry)
	visited[beginWord] = &entry{
		last:  []string{},
		count: 1,
	}
	queue := make([]*bfsContext, 1, 100)
	queue[0] = &bfsContext{
		word:  beginWord,
		count: 1,
	}
	shortest := 0

	for len(queue) > 0 {
		ctx := queue[0]
		queue = queue[1:]

		buf := []byte(ctx.word)
		for i := 0; i < len(ctx.word); i++ {
			ori := buf[i]
			buf[i] = '*'
			key := string(buf)
			next := wordMap[key]
			for i := 0; i < len(next); i++ {
				// 标记是否把next[i]加入到队列中。如果next[i]已经在另一端访问过，则不需要
				if visited[next[i]] == nil {
					visited[next[i]] = &entry{
						last:  []string{ctx.word},
						count: ctx.count + 1,
					}
					queue = append(queue, &bfsContext{
						word:  next[i],
						count: ctx.count + 1,
					})
					if shortest == 0 && next[i] == endWord {
						shortest = ctx.count + 1
					}
				} else {
					// 这个单词之前访问过，但是看看有没有同样短的路径访问这里。更短的路径理论上不会有，因为这是BFS
					et := visited[next[i]]
					if ctx.count+1 == et.count {
						et.last = append(et.last, ctx.word)
					}
				}
			}
			buf[i] = ori
		}
	}

	// DFS连接路径
	if visited[endWord] == nil {
		return [][]string{}
	} else {
		var dfs func(word string) [][]string
		dfs = func(word string) [][]string {
			if len(visited[word].last) == 0 {
				return [][]string{{word}}
			}

			res := make([][]string, 0, 10)
			et := visited[word]
			for i := 0; i < len(et.last); i++ {
				paths := dfs(et.last[i])
				for j := 0; j < len(paths); j++ {
					res = append(res, append(paths[j], word))
				}
			}
			return res
		}
		paths := dfs(endWord)

		return paths
	}

}

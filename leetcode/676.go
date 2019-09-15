package leetcode

type MagicDictionary struct {
	isEnd bool
	next  []*MagicDictionary
}

type my676Node struct {
	// 前面的匹配是否是精确的
	exactMatch bool
	// 前面已经匹配的字符个数（包含可能错误的）
	matched int
	// 即将用于匹配的节点
	node *MagicDictionary
}

/** Initialize your data structure here. */
func Constructor676() MagicDictionary {
	return MagicDictionary{
		isEnd: false,
		next:  make([]*MagicDictionary, 27),
	}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for i := 0; i < len(dict); i++ {
		word := dict[i]

		cur := this
		for j := 0; j < len(word)-1; j++ {
			key := word[j] & 31
			if cur.next[key] == nil {
				cur.next[key] = &MagicDictionary{
					isEnd: false,
					next:  make([]*MagicDictionary, 27),
				}
			}
			cur = cur.next[key]
		}

		cur.next[word[len(word)-1]&31] = &MagicDictionary{
			isEnd: true,
			next:  nil,
		}
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	queue := make([]*my676Node, 1, 16)
	queue[0] = &my676Node{
		exactMatch: true,
		matched:    0,
		node:       this,
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		// 如果已经都matched，则可以返回了
		if cur.matched == len(word) {
			// 仅在词语结束，且非精确的时候才返回true
			if cur.node.isEnd && !cur.exactMatch {
				return true
			}
			continue
		}
		key := word[cur.matched] & 31
		// 采取两种匹配策略
		if cur.exactMatch {
			// 前面全部有精确匹配，但是匹配了一个词语，但是这个串还没有搜索完，因此本次匹配失败，并继续
			if cur.node.isEnd {
				continue
			}
			// 如果前面都是精确的匹配，则可以将本字符的不精确匹配也加入到队列
			for i := 1; i < 27; i++ {
				if cur.node.next[i] != nil {
					queue = append(queue, &my676Node{
						exactMatch: int(key) == i,
						matched:    cur.matched + 1,
						node:       cur.node.next[i],
					})
				}
			}
		} else {
			// 如果前面不是精确的匹配，意味着已经使用了1个字符的错误限制，这里就需要是绝对匹配才行
			if cur.node.next[key] != nil {
				queue = append(queue, &my676Node{
					exactMatch: false,
					matched:    cur.matched + 1,
					node:       cur.node.next[key],
				})
			}
		}
	}
	return false
}

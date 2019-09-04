package leetcode

import "strings"

type my648Trie struct {
	dict map[uint8]*my648Trie
	word *string
}

func newMy648Trie() *my648Trie {
	return &my648Trie{
		dict: make(map[uint8]*my648Trie),
		word: nil,
	}
}

func (trie *my648Trie) insert(word string) {
	cur := trie
	for i := 0; i < len(word); i++ {
		char := word[i]
		node, ok := cur.dict[char]
		if !ok {
			node = &my648Trie{
				dict: make(map[uint8]*my648Trie),
				word: nil,
			}
			cur.dict[char] = node
		}
		cur = node
	}
	cur.word = &word
}

func (trie *my648Trie) find(word string) string {
	cur := trie
	var foundWord *string
	for i := 0; i < len(word); i++ {
		if cur.word != nil {
			foundWord = cur.word
			break
		}
		char := word[i]
		node, ok := cur.dict[char]
		if ok {
			cur = node
		} else {
			// 如果没有了的话，就结束寻找
			break
		}
	}
	if foundWord == nil {
		return word
	} else {
		return *foundWord
	}
}

func replaceWords(dict []string, sentence string) string {
	trie := newMy648Trie()
	for i := 0; i < len(dict); i++ {
		trie.insert(dict[i])
	}

	words := strings.Split(sentence, " ")
	result := ""
	for i := 0; i < len(words); i++ {
		word := trie.find(words[i])
		result += word + " "
	}
	return result[:len(result)-1]
}

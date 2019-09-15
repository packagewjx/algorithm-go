package leetcode

import "testing"

func Test676(t *testing.T) {
	magicDictionary := Constructor676()
	magicDictionary.BuildDict([]string{"hello", "leetcode"})
	println(magicDictionary.Search("hello"))
	println(magicDictionary.Search("hhllo"))
	println(magicDictionary.Search("hell"))
	println(magicDictionary.Search("leetcoded"))

}

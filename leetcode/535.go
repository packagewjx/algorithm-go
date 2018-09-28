package leetcode

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/packagewjx/algorithm-go/util"
)

const BASEURL = "tinyurl.com"
const TINY_LEN = 6

type TinyURL struct {
	urlMap map[string]string
}

func (tiny *TinyURL) init() {
	tiny.urlMap = make(map[string]string)
}

func (tiny *TinyURL) Encode(s string) string {
	hash := calHash(s)

	//我们从hash的0位开始，取出6个字符作为smallHash，存进map中
	//若存在相同的smallHash，则这个值加6，再次取smallHash
	//若整个hash都使用完，则我们将网址和哈希合起来再次哈希，重复这个过程
	splitStart := 0
	original := s
	smallHash := hash[0:TINY_LEN]
	_, ok := tiny.urlMap[smallHash]
	//若ok是true，代表Map已经有相同哈希的网址，我们需要更换val
	for ok {
		if splitStart+TINY_LEN > len(hash) {
			s = s + hash
			hash = calHash(s)
		}
		splitStart += TINY_LEN
		smallHash = hash[splitStart : splitStart+TINY_LEN]
		_, ok = tiny.urlMap[smallHash]
	}
	tiny.urlMap[smallHash] = original
	return BASEURL + "/" + smallHash
}

func (tiny *TinyURL) Decode(s string) string {
	smallHash := s[len(s)-TINY_LEN:]
	return tiny.urlMap[smallHash]
}

//Calculate the sha1 hash for a string
func calHash(s string) string {
	hasher := sha1.New()

	bytes := util.MString(s).ToBytes()
	hasher.Write(bytes)
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

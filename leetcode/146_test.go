package leetcode

import "testing"

func Test146(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)
	cache.Get(2)    // 返回  1
	cache.Put(3, 2) // 该操作会使得密钥 2 作废
	cache.Get(2)    // 返回 -1 (未找到)
	cache.Get(3)    // 返回  3
}

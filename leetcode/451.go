package leetcode

import "sort"

func frequencySort(s string) string {
	type temp struct {
		char  uint8
		count int
	}

	count := make([]*temp, 128)
	for i := 0; i < 128; i++ {
		count[i] = &temp{
			char:  uint8(i),
			count: 0,
		}
	}

	for i := 0; i < len(s); i++ {
		count[s[i]].count++
	}

	sort.Slice(count, func(i, j int) bool {
		return count[i].count > count[j].count
	})

	buf := make([]byte, len(s))
	index := 0
	for i := 0; i < len(count) && count[i].count > 0; i++ {
		for j := 0; j < count[i].count; j++ {
			buf[index] = count[i].char
			index++
		}
	}

	return string(buf)
}

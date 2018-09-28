package leetcode

type stat struct {
	// 本字符的开始位置
	start int
	//
	end int
}

func partitionLabels(S string) []int {
	// 首先做一个统计，统计字母开始和结束的位置
	// 这个数组统计字母最开始出现的顺序
	appears := make([]int32, 0, 26)
	stats := make(map[int32]*stat, 26)
	result := make([]int, 0)
	for i, char := range S {
		s, ok := stats[char]
		if ok {
			s.end = i
		} else {
			stats[char] = &stat{start: i, end: i}
			appears = append(appears, char)
		}
	}

	// 首先假设第一个字母就能分区
	partEnd := stats[appears[0]].end
	partStart := 0
	for i := 1; i < len(appears); i++ {
		//逐个查看这个字母能否包含在内
		char := stats[appears[i]]
		if char.start > partEnd {
			//如果当前字母的开始，比partEnd要远，说明不在这个part里面了，前面的已经包含了仅在该区域出现的字母了
			result = append(result, partEnd-partStart+1)
			partEnd = char.end
			partStart = char.start
		} else if char.end > partEnd {
			//    如果这个字母的结束比当前的end要远，则更新end
			partEnd = char.end
		}
		//这里就是包含在这个part的
	}
	result = append(result, len(S)-partStart)
	return result
}

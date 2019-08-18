package leetcode

// 8皇后问题的变式
func countArrangementRecursive(cur int, choice map[int][]int, remain map[int]bool, result *int) {
	num := choice[cur]
	for i := 0; i < len(num); i++ {
		if remain[num[i]] {
			if cur == len(choice) {
				// 这是最后一个数字
				*result++
			} else {
				// 不是最后一个数字的话，就进入下一层
				remain[num[i]] = false
				countArrangementRecursive(cur+1, choice, remain, result)
				remain[num[i]] = true
			}
		}
	}
}

func countArrangement(N int) int {
	// 记录下标能放置的数字
	choice := make(map[int][]int)
	// 记录什么数字还存在
	remain := make(map[int]bool)
	for i := 1; i <= N; i++ {
		remain[i] = true
		choice[i] = append(choice[i], i)
		for cheng := i << 1; cheng <= N; cheng += i {
			choice[cheng] = append(choice[cheng], i)
			choice[i] = append(choice[i], cheng)
		}
	}
	result := 0
	countArrangementRecursive(1, choice, remain, &result)
	return result
}

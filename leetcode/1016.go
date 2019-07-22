package leetcode

func HighestOne(N int) int {
	if N == 0 {
		return -1
	}
	move := uint(16)
	ret := uint(0)
	for move > 0 {
		if N>>move != 0 {
			ret += move
			N = N >> move
		}
		move >>= 1
	}
	return int(ret)
}

func toInt(binString string) int {
	result := 0
	for i := 0; i < len(binString); i++ {
		result = (result << 1) | int(binString[i]-'0')
	}
	return result
}

func querySpecificLength(S string, length, N int, onePos []int) bool {
	base := 1 << (uint(length) - 1)
	numNeeded := N - base + 1
	// 若需要的数字，比1的总数少，则不可能足够
	if numNeeded > len(onePos) {
		return false
	}

	exists := make(map[int]bool)
	maximumNumber := 1 << (uint(length) - 1)
	for i := 0; i < len(onePos) && onePos[i]+length <= len(S); i++ {
		subStr := S[onePos[i] : onePos[i]+length]
		exists[toInt(subStr)] = true
		if len(exists) == maximumNumber {
			// 如果所有数字均在map中可以提前返回
			return true
		}
	}
	for i := base; i <= N; i++ {
		if val, _ := exists[i]; !val {
			return false
		}
	}
	return true

}

func queryString(S string, N int) bool {
	if N > 2048 {
		return false
	}

	// 有多少个1，就可能有多少个以这个1开头的，长度与N相同的数字，虽然不可能达到这个最大值
	onePos := make([]int, 0, len(S))
	for i := 0; i < len(S); i++ {
		if S[i]-'0' == 1 {
			onePos = append(onePos, i)
		}
	}

	highOnePos := HighestOne(N)
	// 判断最长的是否足够数量
	if !querySpecificLength(S, highOnePos+1, N, onePos) {
		return false
	}

	// 从N小1位的数字开始查看
	for i := highOnePos; i >= 1; i-- {
		max := 1<<uint(i) - 1
		if !querySpecificLength(S, i, max, onePos) {
			return false
		}
	}

	return true
}

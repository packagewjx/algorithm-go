package leetcode

func balancedString(s string) int {
	target := len(s) / 4
	// p站在替换的前端（包含），q站在替换字符串的后端（不包含）
	p := 0
	q := len(s)
	count := make(map[uint8]int)

	// 首先，尽可能让p与q靠近
	for p < len(s) && count[s[p]] < target {
		count[s[p]]++
		p++
	}
	if p == len(s) {
		// 已经平衡，无需继续
		return 0
	}
	// 这里q一定大于p
	for count[s[q-1]] < target {
		count[s[q-1]]++
		q--
	}
	// 最小初始化为当前窗口的大小
	min := q - p
	// 从这个窗口开始，两方向探测。首先存档
	P := p
	Q := q

	// 往p与q减少的方向探测
	for p >= 0 {
		// q走一步，查看是否pq之外的字符串超过了target，若是，则让p回退，若否，则继续
		for p > 0 && count[s[q-1]] >= target {
			// p 回退
			p--
			count[s[p]]--
		}
		if count[s[q-1]] >= target {
			// 若还是大于target，说明q不能再后退了
			break
		}
		q--
		count[s[q]]++

		// 判断当前长度是否更短
		if q-p < min {
			min = q - p
		}
	}

	p = P
	q = Q
	// 往p与q增加的方向探测
	for q <= len(s) {
		// q回退
		for q < len(s) && count[s[p]] >= target {
			count[s[q]]--
			q++
		}
		if count[s[p]] >= target {
			break
		}
		count[s[p]]++
		p++

		if q-p < min {
			min = q - p
		}
	}

	return min
}

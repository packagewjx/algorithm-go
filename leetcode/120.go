package leetcode

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	var buf [2][]int
	buf[0] = make([]int, 0, len(triangle))
	buf[1] = make([]int, 0, len(triangle))

	// 将最后一行复制
	for i := 0; i < len(triangle); i++ {
		buf[(len(triangle)-1)&1] = append(buf[(len(triangle)-1)&1], triangle[len(triangle)-1][i])
	}

	// 从倒数第二行开始查看
	for i := len(triangle) - 2; i >= 0; i-- {
		dp := &buf[i&1]
		*dp = (*dp)[:0]
		next := &buf[i&1^1]

		for j := 0; j < len(triangle[i]); j++ {
			if (*next)[j] < (*next)[j+1] {
				*dp = append(*dp, triangle[i][j]+(*next)[j])
			} else {
				*dp = append(*dp, triangle[i][j]+(*next)[j+1])
			}
		}
	}

	return buf[0][0]
}

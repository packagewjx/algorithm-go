package leetcode

import (
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	jiecheng := make([]int, n)
	jiecheng[0] = 1
	for i := 1; i < n; i++ {
		jiecheng[i] = jiecheng[i-1] * (i + 1)
	}

	res := strings.Builder{}
	// 变成从0开始计数
	k--
	for i := n - 1; i > 0; i-- {
		pos := k / jiecheng[i-1]
		res.WriteString(strconv.Itoa(nums[pos]))
		// 删除pos的数
		nums = append(nums[:pos], nums[pos+1:]...)
		k = k % jiecheng[i-1]
	}
	res.WriteByte(byte(0x30 | nums[0]))

	return res.String()
}

package leetcode

func removeDuplicates(s string, k int) string {
	if k == 0 {
		return s
	}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) >= k {
			c := stack[len(stack)-1]
			match := true
			for j := len(stack) - 2; j >= len(stack)-k; j-- {
				if stack[j] != c {
					match = false
					break
				}
			}
			if match {
				// 删除栈顶的k元素
				stack = stack[:len(stack)-k]
			}
		}
	}
	return string(stack)
}

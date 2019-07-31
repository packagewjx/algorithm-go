package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, 10)
	exist := make(map[int]bool)
	deleted := make(map[int]bool)
	pushedCnt := 0
	for i := 0; i < len(popped); i++ {
		if deleted[popped[i]] {
			continue
		}
		pop := popped[i]
		for ; pushedCnt < len(pushed) && pushed[pushedCnt] != pop; pushedCnt++ {
			stack = append(stack, pushed[pushedCnt])
			exist[pushed[pushedCnt]] = true
		}
		pushedCnt++
		deleted[pop] = true
		for j := i + 1; len(stack) > 0 && j < len(popped); j++ {
			if exist[popped[j]] && stack[len(stack)-1] == popped[j] {
				stack = stack[0 : len(stack)-1]
				exist[popped[j]] = false
				deleted[popped[j]] = true
			}
		}
		// 到这里应该为空
		if len(stack) != 0 {
			return false
		}
	}
	return true
}

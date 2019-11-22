package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, 10)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "*":
			stack[len(stack)-2] = stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "+":
			stack[len(stack)-2] = stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] = stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] = stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		}
	}

	return stack[0]
}

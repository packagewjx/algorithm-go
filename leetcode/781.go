package leetcode

func numRabbits(answers []int) int {
	numMap := make(map[int]int)
	for i := 0; i < len(answers); i++ {
		numMap[answers[i]] = numMap[answers[i]] + 1
	}
	result := 0
	for num, count := range numMap {
		result += count / (num + 1) * (num + 1)
		if count%(num+1) != 0 {
			result += num + 1
		}
	}
	return result
}

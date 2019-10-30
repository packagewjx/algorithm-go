package leetcode

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	for _, v := range nums {
		numMap[v] = true
	}

	max := 0
	for num := range numMap {
		if numMap[num-1] {
			continue
		}

		streak := 1
		cur := num + 1
		for numMap[cur] {
			streak++
			cur++
		}

		if streak > max {
			max = streak
		}
	}
	return max
}

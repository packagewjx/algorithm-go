package leetcode

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	countCount := make(map[int]bool)
	for _, c := range count {
		if countCount[c] {
			return false
		}
		countCount[c] = true
	}

	return true
}

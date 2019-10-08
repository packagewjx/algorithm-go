package leetcode

func minCostToMoveChips(chips []int) int {
	oddCount := 0
	evenCount := 0
	for i := 0; i < len(chips); i++ {
		if chips[i]&1 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	if oddCount > evenCount {
		return evenCount
	} else {
		return oddCount
	}
}

package leetcode

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	smallest := math.MaxInt64
	smPair := make([][]int, 0)
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < smallest {
			smallest = diff
			smPair = smPair[:0]
			smPair = append(smPair, []int{arr[i], arr[i+1]})
		} else if diff == smallest {
			smPair = append(smPair, []int{arr[i], arr[i+1]})
		}
	}
	return smPair
}

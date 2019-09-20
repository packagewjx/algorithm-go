package leetcode

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0, len(intervals))
	end := math.MinInt64
	for i := 0; i < len(intervals); i++ {
		if i == 0 || intervals[i][0] > end {
			res = append(res, intervals[i])
			end = intervals[i][1]
		} else if intervals[i][0] <= end {
			// 这里i一定大于1
			if intervals[i][1] > end {
				// 合并
				res[len(res)-1][1] = intervals[i][1]
				end = intervals[i][1]
			}
		}
	}
	return res
}

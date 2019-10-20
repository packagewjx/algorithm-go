package leetcode

import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	type job struct {
		start  int
		end    int
		profit int
	}
	jobs := make([]*job, len(startTime))
	for i := 0; i < len(startTime); i++ {
		jobs[i] = &job{
			start:  startTime[i],
			end:    endTime[i],
			profit: profit[i],
		}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].start < jobs[j].start
	})

	dp := make([]int, len(jobs))
	dp[len(jobs)-1] = jobs[len(jobs)-1].profit
	for i := len(jobs) - 2; i >= 0; i-- {
		max := 0
		// 接受本工作，则查看结束时间的最大值
		pos := sort.Search(len(jobs), func(j int) bool {
			return jobs[j].start >= jobs[i].end
		})
		max = jobs[i].profit
		if pos != len(startTime) {
			max += dp[pos]
		}
		// 不接受本工作，则查看下一个工作的最大值
		if dp[i+1] > max {
			max = dp[i+1]
		}
		dp[i] = max
	}

	return dp[0]
}

package leetcode

import "math"

var dayPass = []int{1, 7, 30}

func mincostTicketsDP(days []int, costs []int, start int, memo *map[int]int) int {
	if start >= len(days) {
		return 0
	}
	cost, ok := (*memo)[start]
	if ok {
		return cost
	}

	minCost := math.MaxInt64
	for i := 0; i < len(dayPass); i++ {
		lastDay := days[start] + dayPass[i] - 1
		coverUntil := start + 1
		for ; coverUntil < len(days); coverUntil++ {
			if days[coverUntil] > lastDay {
				break
			}
		}
		cost = costs[i] + mincostTicketsDP(days, costs, coverUntil, memo)
		if cost < minCost {
			minCost = cost
		}
	}
	(*memo)[start] = minCost
	return minCost
}

func mincostTickets(days []int, costs []int) int {
	memo := make(map[int]int, len(days))
	return mincostTicketsDP(days, costs, 0, &memo)
}

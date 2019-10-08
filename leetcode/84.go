package leetcode

func largestRectangleArea(heights []int) int {
	max := 0
	for i := 0; i < len(heights); i++ {
		min := heights[i]
		for j := i; j < len(heights); j++ {
			if min > heights[j] {
				min = heights[j]
			}
			area := min * (j - i + 1)
			if area > max {
				max = area
			}
		}
	}

	return max
}

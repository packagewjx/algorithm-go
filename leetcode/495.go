package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	sum := 0
	end := timeSeries[0] + duration
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i] < end {
			sum += timeSeries[i] - timeSeries[i-1]
		} else {
			sum += duration
		}
		end = timeSeries[i] + duration
	}
	return sum + end - timeSeries[len(timeSeries)-1]
}

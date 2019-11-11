package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	//参考官方题解得到的代码
	totalTank := 0
	currTank := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		totalTank += gas[i] - cost[i]
		currTank += gas[i] - cost[i]
		if currTank < 0 {
			start = i + 1
			currTank = 0
		}
	}

	if totalTank < 0 {
		return -1
	} else {
		return start
	}
}

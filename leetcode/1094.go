package leetcode

func carPooling(trips [][]int, capacity int) bool {
	passengersOfKM := make([]int, 1000)
	for i := 0; i < len(trips); i++ {
		numOfPassengers := trips[i][0]
		start := trips[i][1]
		end := trips[i][2]

		for j := start; j < end; j++ {
			passengersOfKM[j] = passengersOfKM[j] + numOfPassengers
			if passengersOfKM[j] > capacity {
				return false
			}
		}
	}

	return true
}

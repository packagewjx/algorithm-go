package contest

import "sort"

func robot(command string, obstacles [][]int, x int, y int) bool {
	height := 0
	width := 0
	for i := 0; i < len(command); i++ {
		if command[i] == 'u' {
			height++
		} else {
			width++
		}
	}

	canGo := make([][]bool, height)
	for i := 0; i < len(canGo); i++ {
		canGo[i] = make([]bool, width)
	}

	sort.Slice(obstacles, func(i, j int) bool {
		return obstacles[i][0]+obstacles[i][1] < obstacles[j][0]+obstacles[j][1]
	})
	obsIndex := 0

	upperBound := x + y
	for i := 0; i*(height+width) <= upperBound; i++ {
		for ; obstacles[obsIndex][0]+obstacles[obsIndex][1] <= i*(height+width); obsIndex++ {
			ax := obstacles[obsIndex][0] - i*width
			ay := obstacles[obsIndex][1] - i*height
			if command[0] == 'U' {
				ay -= 1
			} else {
				ax -= 1
			}
			if canGo[ay][ax] {
				return false
			}
		}

	}

}

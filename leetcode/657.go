package leetcode

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, action := range moves {
		switch action {
		case 'U':
			y += 1
		case 'D':
			y -= 1
		case 'L':
			x -= 1
		case 'R':
			x += 1
		}
	}

	return x == 0 && y == 0
}

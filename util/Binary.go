package util

func NumOfOnes(num int) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

// 计算整数中最高位1的位置，以0为起始位置。如果是0，则返回-1
func HighestOne(N int) int {
	if N == 0 {
		return -1
	}
	move := uint(16)
	ret := uint(0)
	for move > 0 {
		if N>>move != 0 {
			ret += move
			N = N >> move
		}
		move >>= 1
	}
	return int(ret)
}

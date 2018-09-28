package util

func NumOfOnes(num int) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

package leetcode

import "github.com/packagewjx/algorithm-go/util"

func main() {
	print(util.NumOfOnes(8))
}

func hammingDistance(x int, y int) int {
	return NumOfOnes(x ^ y)
}

func NumOfOnes(num int) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

package leetcode

func candy(ratings []int) int {
	candies := make([]int, len(ratings))

	// 正向遍历
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	// 反向遍历
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	// 首先每人至少一个
	sum := len(ratings) * 1
	// candies里面就是加的数量
	for i := 0; i < len(candies); i++ {
		sum += candies[i]
	}

	return sum
}

package leetcode

func intToRoman(num int) string {
	nums := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	cur := len(romans) - 1
	result := ""
	for num > 0 {
		for nums[cur] > num {
			cur--
		}
		num -= nums[cur]
		result += romans[cur]
	}
	return result
}

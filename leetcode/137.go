package leetcode

/* 总体思想是，设计一个遇见3次1就归零的布尔表达式
00 运算 1 = ab
ab 运算 1 = cd
cd 运算 1 = 00
*/
func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, x := range nums {
		b = (b ^ x) & ^a
		a = (a ^ x) & ^b
	}

	return b
}

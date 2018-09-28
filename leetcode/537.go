package leetcode

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	as := strings.Split(a, "+")
	bs := strings.Split(b, "+")
	num1, _ := strconv.Atoi(as[0])
	num2, _ := strconv.Atoi(as[1][:len(as[1])-1])
	num3, _ := strconv.Atoi(bs[0])
	num4, _ := strconv.Atoi(bs[1][:len(bs[1])-1])
	r1 := num1*num3 - num2*num4
	r2 := num1*num4 + num2*num3
	return strconv.Itoa(r1) + "+" + strconv.Itoa(r2) + "i"
}

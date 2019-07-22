package leetcode

import "testing"

func Test10156(t *testing.T) {
	queryString("1011100101000", 12)
}

func TestToInt(t *testing.T) {
	print(toInt("101"), "\n")
	print(toInt("1"), "\n")
	print(toInt("0"), "\n")
	print(toInt("10"), "\n")
	print(toInt("110"), "\n")
}

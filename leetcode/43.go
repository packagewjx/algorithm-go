package leetcode

// 竖式乘法
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	} else if num1 == "1" {
		return num2
	} else if num2 == "1" {
		return num1
	}

	buf := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			pos := i + j + 1
			cheng := (num1[i] & 15) * (num2[j] & 15)
			n1 := cheng % 10
			jinwei := cheng / 10
			buf[pos] += n1
			if buf[pos] >= 10 {
				buf[pos] -= 10
				jinwei += 1
			}
			// 循环进位
			pos--
			for jinwei > 0 {
				buf[pos] += jinwei
				if buf[pos] >= 10 {
					buf[pos] -= 10
					jinwei = 1
				} else {
					jinwei = 0
				}
				pos--
			}

		}
	}
	var start int
	for start = 0; start < len(buf) && buf[start] == 0; start++ {
	}
	for i := start; i < len(buf); i++ {
		buf[i] |= 0x30
	}

	return string(buf[start:])
}

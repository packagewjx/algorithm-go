package study

import (
	"encoding/binary"
)

// 试着以32位的形式处理由小写转换大写的

// 函数将字节数组转换为int数组且不变换顺序
func byteArrayToIntArray(array []byte) []uint32 {
	nums := make([]uint32, 0, len(array)>>2+1)
	for len(array) > 4 {
		nums = append(nums, binary.BigEndian.Uint32(array[:4]))
		array = array[4:]
	}
	shift := uint32(24)
	lastNum := uint32(0)
	for i := 0; i < len(array); i++ {
		lastNum = lastNum | uint32(array[i])<<shift
		shift -= 8
	}
	nums = append(nums, lastNum)
	return nums
}

// uint32数组转换成byte数组，去掉最后的0，假定0出现在最后一个uint32上
func uint32ArrayToByteArrayWithNoLast0(array []uint32) []byte {
	b := make([]byte, 0, len(array)<<2)
	for i := 0; i < len(array)-1; i++ {
		b = append(b, byte(array[i]>>24))
		b = append(b, byte(array[i]>>16))
		b = append(b, byte(array[i]>>8))
		b = append(b, byte(array[i]))
	}

	shift := uint32(24)
	num := byte(array[len(array)-1] >> shift)

	for num != 0 {
		b = append(b, num)
		shift -= 8
		num = byte(array[len(array)-1] >> shift)
	}
	return b
}

func toUpperCase(text string) string {
	intArray := byteArrayToIntArray([]byte(text))
	for key, val := range intArray {
		result := uint32toUpperCase(val)
		intArray[key] = result
	}
	newBytes := uint32ArrayToByteArrayWithNoLast0(intArray)
	return string(newBytes)
}

// 核心算法
func uint32toUpperCase(char uint32) uint32 {
	// 首先查看应该所有字节都大于等于97，也就是每一字节减去97应该为正数
	// 也就是加上-97给每一位
	num1 := char + 0x9F9F9F9F
	// 取每字节最高位，然后取反
	smallerThan97 := num1&0x80808080 ^ 0x80808080
	// 减去进位
	num1 -= smallerThan97 << 1
	smallerThan97 = num1 & 0x80808080

	// 全部小于，立即返回
	if smallerThan97 == 0x80808080 {
		return char
	}

	// 是否小于等于122，也就是小于123，也就是减去123应该是负数，也就是加上-123
	num1 = char + 0x85858585
	biggerThan122 := num1&0x80808080 ^ 0x80808080

	// 若两个字节相减，得到正数，将会有一个进位在第9个比特位置，虽然对字节没影响
	// 但是会影响我们这个方法，给前一个字节加了1，因此需要减掉，
	// 只需左移biggerThan123然后减去就行吧
	num1 -= biggerThan122 << 1
	biggerThan122 = num1 & 0x80808080

	// 全部大于，立即返回
	if biggerThan122 == 0 {
		return char
	}

	// 若在小数区间内，应该减去的数，但是有些不是，因此需要裁剪这个数
	shouldDecrease := uint32(0x20202020)

	// 把smallerThan97和biggerThan123右移两位相减即可，因为32只有一个bit，
	// 恰好减了就不用加这个数
	// 真聪明啊a和A间隔32
	// 最高位为1代表小于97，因此是异或运算，除去这个1
	shouldDecrease = shouldDecrease ^ (smallerThan97 >> 2)
	// 因为最高位为1代表小于123，因此是与运算
	shouldDecrease = shouldDecrease & (biggerThan122 >> 2)

	return char - shouldDecrease
}
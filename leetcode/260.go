package leetcode

/**
一个非常聪明的办法。首先还是所有的数字异或一次，得出来是两个数字（称为p，q）的异或，那么这个结果中的1，代表着p和q不同的位。
如果我们将这一个位置上为0的数异或到一起，为1的数异或到一起，那么就会分成两堆数，而因为除了p，q之外的数都是重复两次的，那么这两个结果就是p和q了
*/
func singleNumber(nums []int) []int {
	a := 0
	for i := 0; i < len(nums); i++ {
		a ^= nums[i]
	}

	mask := uint(1)
	for a&1 != 1 {
		a >>= 1
		mask <<= 1
	}
	umask := uint(0xFFFFFFFFFFFFFFFF) ^ mask

	p := 0
	q := 0
	for i := 0; i < len(nums); i++ {
		if uint(nums[i])&mask == mask {
			p ^= nums[i]
			continue
		}
		if uint(nums[i])|umask == umask {
			q ^= nums[i]
		}
	}

	return []int{p, q}
}

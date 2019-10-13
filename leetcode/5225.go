package leetcode

func maxEqualFreq(nums []int) int {
	count := make(map[int]int)
	countNum := make(map[int]map[int]bool)

	max := 0

	for i := 0; i < len(nums); i++ {
		originalCount := count[nums[i]]
		count[nums[i]]++
		// 从本来的里面删除
		if countNum[originalCount] != nil {
			delete(countNum[originalCount], nums[i])
			if len(countNum[originalCount]) == 0 {
				delete(countNum, originalCount)
			}
		}
		if countNum[count[nums[i]]] == nil {
			countNum[count[nums[i]]] = make(map[int]bool)
		}
		countNum[count[nums[i]]][nums[i]] = true

		if len(countNum) == 2 {
			// 说明目前有每个数字的总数，有2个取值
			keys := make([]int, 0, 2)
			for count := range countNum {
				keys = append(keys, count)
			}

			// 多了一个数出来
			if len(countNum[keys[0]]) == 1 && keys[0]-keys[1] == 1 {
				max = i + 1
			}
			if len(countNum[keys[1]]) == 1 && keys[1]-keys[0] == 1 {
				max = i + 1
			}
			// 如果删掉这个数，就没了这个数，剩下其他的数都是出现频率相同
			if len(countNum[1]) == 1 {
				max = i + 1
			}
		} else if len(countNum) == 1 {
			// 代表着所有数字都只出现了1次，因此可以删除随便一个得到符合的前缀
			for count := range countNum {
				if count == 1 {
					max = i + 1
				}
			}
		}
	}

	if len(count) == 1 {
		// 如果只有1个数字，则就是最大了
		max = len(nums)
	}

	return max
}

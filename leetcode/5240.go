package leetcode

func maxLengthBT(arr []string, cur int, curMap map[byte]bool, pos int, max *int) {
	for i := pos; i < len(arr); i++ {
		noExist := true
		for j := 0; j < len(arr[i]); j++ {
			if curMap[arr[i][j]] {
				noExist = false
			}
		}
		if noExist {
			if cur+len(arr[i]) > *max {
				*max = cur + len(arr[i])
			}

			// 进入下一个循环
			for j := 0; j < len(arr[i]); j++ {
				curMap[arr[i][j]] = true
			}
			maxLengthBT(arr, cur+len(arr[i]), curMap, i+1, max)
			for j := 0; j < len(arr[i]); j++ {
				curMap[arr[i][j]] = false
			}
		}
	}
}

func maxLength(arr []string) int {
	max := 0

	// 检查字符串本身是否超过
	for i := 0; i < len(arr); i++ {
		arrMap := make(map[byte]bool)
		for j := 0; j < len(arr[i]); j++ {
			if arrMap[arr[i][j]] {
				// 删除这条字符串
				arr = append(arr[:i], arr[i+1:]...)
				i--
				break
			}
			arrMap[arr[i][j]] = true
		}
	}

	maxLengthBT(arr, 0, map[byte]bool{}, 0, &max)
	return max
}

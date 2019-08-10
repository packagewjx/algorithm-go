package util

/**
二分搜索。如果找到了key，则返回其index，如果没有，则返回适合插入位置的index的-(index+1)，该位置将会是第一个大于这个数的数。
算法永远试图寻找最左边的相等值。
*/
func BinarySearch(arr []int, key int) (index int) {
	begin := 0
	end := len(arr)
	// 目的是在偶数长度时，找到中间两个数的较小的值
	cur := begin + (end-begin)/2 - (1 - (end-begin)&1)
	for begin+1 < end {
		if arr[cur] > key {
			end = cur
		} else if arr[cur] < key {
			begin = cur + 1
		} else /*arr[cur] == key*/ {
			end = cur + 1
		}
		cur = begin + (end-begin)/2 - (1 - (end-begin)&1)
	}

	if arr[cur] == key {
		return cur
	} else {
		if arr[cur] > key {
			return -(cur + 1)
		} else {
			// 返回第一个key大的位置
			return -(cur + 2)
		}
	}
}

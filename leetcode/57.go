package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	// 二分查找
	l := 0
	r := len(intervals)
	for l < r {
		mid := (l + r) / 2
		if intervals[mid][0] > newInterval[0] {
			r = mid
		} else if intervals[mid][1] >= newInterval[0] {
			l = mid
			r = mid
			break
		} else {
			l = mid + 1
		}
	}

	if l == len(intervals) {
		// 落在了最后的位置
		return append(intervals, newInterval)
	} else {
		if intervals[l][0] > newInterval[0] {
			// 头部没有落入区间内
			if intervals[l][0] <= newInterval[1] {
				if intervals[l][1] >= newInterval[1] {
					// 新区间的尾部落入了区间内，合并两个区间然后返回
					intervals[l] = []int{newInterval[0], intervals[l][1]}
					return intervals
				} else /* intervals[l][1] < newInterval[1] */ {
					// 新区间的尾部没有落入本区间，这时需要合并多个
					combine := newInterval
					delCount := 1
					for i := l + 1; i < len(intervals) && intervals[i][0] <= combine[1]; i++ {
						delCount++
						if combine[1] > intervals[i][1] {
							continue
						}
						combine = []int{combine[0], intervals[i][1]}
					}

					return append(append(intervals[:l], combine), intervals[l+delCount:]...)
				}

			} else /*intervals[l][1] > newInterval[1]*/ {
				// 没有的话，直接插入
				if l == 0 {
					return append([][]int{newInterval}, intervals...)
				} else {
					temp := append(intervals[:l], intervals[l-1:]...)
					temp[l] = newInterval
					return temp
				}
			}
		} else /*intervals[l][0] <= newInterval[0] */ {
			// 落入了区间内
			if intervals[l][1] >= newInterval[1] {
				// 包含了整个区间，则无需插入
				return intervals
			} else /* intervals[l][1] < newInterval[1] */ {
				// 合并多个
				combine := []int{intervals[l][0], newInterval[1]}
				delCount := 1

				for i := l + 1; i < len(intervals) && intervals[i][0] <= combine[1]; i++ {
					delCount++
					if combine[1] > intervals[i][1] {
						continue
					}
					combine = []int{combine[0], intervals[i][1]}
				}

				return append(append(intervals[:l], combine), intervals[l+delCount:]...)
			}
		}
	}

}

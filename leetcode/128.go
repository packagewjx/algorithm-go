//+build 128

package leetcode

func longestConsecutive(nums []int) int {
	max := 0

	type entry struct {
		start int
		end   int
	}
	// 放在重复的数字上标识
	visited := &entry{
		start: 0,
		end:   0,
	}
	lenMap := make(map[int]*entry)

	for i := 0; i < len(nums); i++ {
		if lenMap[nums[i]] != nil {
			continue
		}
		left := lenMap[nums[i]-1]
		right := lenMap[nums[i]+1]

		if left == nil {
			if right == nil {
				// 两个都是nil
				lenMap[nums[i]] = &entry{
					start: nums[i],
					end:   nums[i],
				}
			} else {
				// 左边为nil，右边的数不是nil
				// nums[i]成为区间的开始
				lenMap[nums[i]] = &entry{
					start: nums[i],
					end:   right.end,
				}
				// 区间末尾的元素的start改为这个数字
				lenMap[right.end].start = nums[i]
				// 删除掉这个区间中间的元素，因为不可能会查询到
				if right.end != nums[i]+1 {
					// 标记本数字已经访问过，包含在区间内，后面不再继续
					lenMap[nums[i]+1] = visited
				}
			}
		} else /*lenMap[nums[i]-1]!=nil*/ {
			if right != nil {
				// 两边都不是nil，这个数把两边的区间连接起来了
				lenMap[left.start].end = right.end
				lenMap[right.end].start = left.start
				// 非区间边界的所有数字都成visited
				if left.start != nums[i]-1 {
					lenMap[nums[i]-1] = visited
				}
				if right.end != nums[i]+1 {
					lenMap[nums[i]+1] = visited
				}
				// 把自己也要变成visited
				lenMap[nums[i]] = visited
			} else {
				// 左边的不是nil，右边为nil
				lenMap[left.start].end = nums[i]
				lenMap[nums[i]] = &entry{
					start: left.start,
					end:   nums[i],
				}
				// 非区间边界的数字都成visited
				if left.start != nums[i]-1 {
					lenMap[nums[i]-1] = visited
				}
			}
		}
	}

	for _, et := range lenMap {
		if et != visited && et.end-et.start+1 > max {
			max = et.end - et.start + 1
		}
	}

	return max
}

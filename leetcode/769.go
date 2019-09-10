package leetcode

func maxChunksToSorted(arr []int) int {
	count := 0
	for i := 0; i < len(arr); {
		end := arr[i]
		// 这个分区一定会包含从i到end的所有数
		j := i
		for ; j <= end && j < len(arr); j++ {
			if arr[j] > end {
				end = arr[j]
			}
		}
		count++
		i = j
	}
	return count
}

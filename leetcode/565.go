package leetcode

func arrayNesting(nums []int) int {
	visited := make([]bool, len(nums))
	longest := 0
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		set := make([]bool, len(nums))
		count := 0
		cur := i
		for set[cur] == false {
			visited[cur] = true
			count++
			set[cur] = true
			cur = nums[cur]
		}
		if count > longest {
			longest = count
		}
	}
	return longest
}

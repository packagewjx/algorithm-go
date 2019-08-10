package leetcode

// 使用后面一个节点找到的比该节点大的结果，如果比我小，则继续找这个节点的比他大的节点
func nextLargerNodesRecursive(head *ListNode, index int) (nums, biggerThanIndex []int) {
	if head.Next == nil {
		nums = make([]int, index+1)
		biggerThanIndex = make([]int, index+1)
		nums[index] = head.Val
		// -1代表后面没有数字比我大
		biggerThanIndex[index] = -1
		return
	}

	nums, biggerThanIndex = nextLargerNodesRecursive(head.Next, index+1)
	nums[index] = head.Val
	nextIndex := index + 1
	for nextIndex != -1 {
		if nums[nextIndex] > head.Val {
			biggerThanIndex[index] = nextIndex
			// 找到了，立即退出
			return
		}
		nextIndex = biggerThanIndex[nextIndex]
	}
	// 没有找到，因此赋值为-1
	biggerThanIndex[index] = -1

	return
}

func nextLargerNodes(head *ListNode) []int {
	nums, biggerThanIndex := nextLargerNodesRecursive(head, 0)
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if biggerThanIndex[i] != -1 {
			result[i] = nums[biggerThanIndex[i]]
		}
	}
	return result
}

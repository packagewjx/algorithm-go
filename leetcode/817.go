package leetcode

func numComponents(head *ListNode, G []int) int {
	numMap := make(map[int]bool)
	for i := 0; i < len(G); i++ {
		numMap[G[i]] = true
	}

	cur := head
	count := 0
	lastInG := false
	for cur != nil {
		if numMap[cur.Val] {
			if lastInG {
				count++
			}
			lastInG = true
		} else {
			lastInG = false
		}
		cur = cur.Next
	}
	return count + len(numMap) - count*2
}

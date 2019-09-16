package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 先走n-1步
	cur := head
	for i := 0; i < n-1; i++ {
		if cur == nil {
			// 无法删除
			return head
		}
		cur = cur.Next
	}
	// 如果刚好是nil，则返回了
	if cur == nil {
		return head
	}
	// 如果下一个是空的话，说明刚好有n个节点，因此删除head
	if cur.Next == nil {
		ret := head.Next
		head.Next = nil
		return ret
	}
	// 再走一步，使得cur.Next是空的时候，刚好删除deleteNext的下一个节点即可
	cur = cur.Next
	deleteNext := head
	for cur.Next != nil {
		cur = cur.Next
		deleteNext = deleteNext.Next
	}
	node := deleteNext.Next.Next
	deleteNext.Next.Next = nil
	deleteNext.Next = node
	return head
}

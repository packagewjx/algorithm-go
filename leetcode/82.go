package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	fakeHead := &ListNode{Next: head}
	last := fakeHead

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			val := cur.Val
			for cur.Next != nil && cur.Next.Val == val {
				cur = cur.Next
			}
			last.Next = cur.Next
			temp := cur
			cur = cur.Next
			// 断开这个已经被删除的连接
			temp.Next = nil
		} else {
			last = cur
			cur = cur.Next
		}
	}
	return fakeHead.Next
}

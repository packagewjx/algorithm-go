package leetcode

func reorderList(head *ListNode) {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	if length <= 1 {
		return
	}

	stack := make([]*ListNode, 0, length/2+1)
	cur = head
	// 走到后面一半
	for i := 0; i < (length+1)/2; i++ {
		cur = cur.Next
	}
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}

	cur = head
	for i := 0; i < length/2; i++ {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		top.Next = cur.Next
		cur.Next = top
		cur = top.Next
	}
	cur.Next = nil
}

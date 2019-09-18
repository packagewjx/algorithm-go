package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head
	second := head.Next
	next := second.Next
	// 先执行头两个指针的转换
	first.Next = next
	second.Next = first
	// 改变头结点
	head = second
	last := first
	next = first.Next
	for next != nil && next.Next != nil {
		first = next
		second = next.Next
		next = second.Next
		// 转换
		first.Next = next
		second.Next = first
		last.Next = second
		// 更新循环变量
		last = first
		next = first.Next
	}

	return head
}

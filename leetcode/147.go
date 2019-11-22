package leetcode

func insertionSortList(head *ListNode) *ListNode {
	// 没有或只有1个返回即可
	if head == nil || head.Next == nil {
		return head
	}

	// 这是排好序的数字的尾部，处理的下一个数字就是tail.Next
	tail := head
	fakeHead := &ListNode{
		Val:  0,
		Next: head,
	}
	for tail.Next != nil {
		sorted := fakeHead
		if tail.Next.Val > tail.Val {
			// 此时，我们不需要换位，把tail往后移即可
			tail = tail.Next
			continue
		}

		// 停下来时，sorted.Next的位置应是大于等于tail.Next的
		for sorted.Next != tail.Next && sorted.Next.Val < tail.Next.Val {
			sorted = sorted.Next
		}

		// 将tail.Next放置到sorted.Next
		temp := tail.Next
		tail.Next = temp.Next
		temp.Next = sorted.Next
		sorted.Next = temp
	}

	return fakeHead.Next
}

package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	// 首先遍历，查看长度
	length := 1
	for p.Next != nil {
		length++
		p = p.Next

	}
	// 取余，赶走多余的步骤
	k = k % length
	if k == 0 {
		// 若是0，则返回即可
		return head
	}

	p = head
	// 行走length-k-1步，这里一定不会遇到nil
	for i := 0; i < length-k-1; i++ {
		p = p.Next
	}

	// p成了最后一个
	newHead := p.Next
	p.Next = nil
	// 找到链表最后一个，并连接到head
	p = newHead
	for p.Next != nil {
		p = p.Next
	}
	p.Next = head

	return newHead
}

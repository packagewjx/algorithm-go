package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		// 单个节点的情况
		return nil
	}

	slow := head
	fast := head
	var intercept *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			intercept = fast
			break
		}
	}

	if intercept == nil {
		return nil
	}

	// 第二阶段，slow重新出发，fast则是在相遇点出发
	slow = head
	fast = intercept
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

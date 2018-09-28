package leetcode

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	go1Node := head
	go2Node := head.Next

	for {
		if go2Node.Next == nil {
			return go1Node.Next
		}
		go2Node = go2Node.Next
		go1Node = go1Node.Next
		if go2Node.Next == nil {
			return go1Node
		}
		go2Node = go2Node.Next
	}
}

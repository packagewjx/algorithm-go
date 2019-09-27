package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for i := 0; i < k-1; i++ {
		if cur.Next == nil {
			return head
		}
		cur = cur.Next
	}
	nextHead := reverseKGroup(cur.Next, k)
	last := head
	cur = head.Next
	for i := 0; i < k-1; i++ {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	head.Next = nextHead
	return last
}

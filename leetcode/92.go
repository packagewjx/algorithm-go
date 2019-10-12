package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// m==n无需处理
	if m == n {
		return head
	}
	fakeHead := &ListNode{Next: head}
	cur := fakeHead
	// 走m-1步
	for i := 1; i < m; i++ {
		cur = cur.Next
	}
	frontEnd := cur
	reverseEnd := cur.Next

	// 要处理的下一个的下一个是第一个翻转的，如果是nil，无需处理
	if cur.Next.Next == nil {
		return fakeHead.Next
	}
	last := cur.Next
	cur = last.Next
	next := cur.Next
	// 由于length >= n >= m，下面不用判断是否为nil
	for i := 0; i < n-m; i++ {
		cur.Next = last

		last = cur
		// 对最后情况的特殊处理
		if next == nil {
			cur = nil
			break
		} else {
			cur = next
			next = cur.Next
		}
	}
	frontEnd.Next = last
	reverseEnd.Next = cur

	return fakeHead.Next
}

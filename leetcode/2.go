package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	fakeHead := &ListNode{Next: nil, Val: 0}
	last := fakeHead
	p := l1
	q := l2
	for p != nil && q != nil {
		newNode := &ListNode{
			Val:  p.Val + q.Val + carry,
			Next: nil,
		}
		if newNode.Val > 9 {
			newNode.Val -= 10
			carry = 1
		} else {
			carry = 0
		}
		last.Next = newNode
		last = newNode
		p = p.Next
		q = q.Next
	}

	if p == nil && q == nil {
		if carry == 1 {
			last.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
	} else {
		if p == nil {
			p = q
		}
		for p != nil {
			newNode := &ListNode{
				Val:  p.Val + carry,
				Next: nil,
			}
			if newNode.Val == 10 {
				newNode.Val = 0
				carry = 1
			} else {
				carry = 0
			}
			last.Next = newNode
			last = newNode
			p = p.Next
		}
		// 这里都是nil了
		if carry == 1 {
			last.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
	}
	return fakeHead.Next
}

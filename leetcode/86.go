//+build 86

package leetcode

func partition(head *ListNode, x int) *ListNode {
	fakeHead := &ListNode{Next: head}
	// 指向小于x的最后一个元素，当然不包括fakeHead
	frontEnd := fakeHead
	for cur := fakeHead; cur != nil && cur.Next != nil; {
		if cur.Next != frontEnd && cur.Next.Val < x {
			// 要判断的是cur.Next，因此如果已经是frontEnd，代表其已经在前面的列表，无需判断，此时要移动
			tihuan := cur.Next
			cur.Next = cur.Next.Next
			tihuan.Next = frontEnd.Next
			frontEnd.Next = tihuan
			frontEnd = frontEnd.Next
		} else {
			cur = cur.Next
		}
	}
	return fakeHead.Next
}

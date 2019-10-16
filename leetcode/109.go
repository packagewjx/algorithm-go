package leetcode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{
			Val: head.Val,
		}
	}

	// 双指针法，分开两个list
	var last *ListNode
	p := head
	q := head.Next
	for q != nil {
		last = p
		p = p.Next
		q = q.Next
		if q == nil {
			break
		}
		q = q.Next
	}

	// 此时p应该在中间，last应该是p的上一个

	if last == nil {
		// last是nil，说明p没有走，左节点为空，只拿右节点即可
		right := sortedListToBST(p.Next)
		return &TreeNode{
			Val:   p.Val,
			Left:  nil,
			Right: right,
		}
	}

	// last不为空，将last的Next暂时赋值为空，进行递归
	last.Next = nil
	root := &TreeNode{
		Val:   p.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(p.Next),
	}
	// 重新连接
	last.Next = p
	return root
}

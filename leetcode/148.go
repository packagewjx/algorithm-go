package leetcode

func sortListQuickSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 快排分区思想，以head为pivot
	l := head
	rightStart := &ListNode{
		Val:  0,
		Next: nil,
	}
	r := rightStart

	next := head.Next
	for next != nil {
		cur := next
		next = cur.Next
		// 清除，防止产生错误
		cur.Next = nil
		if cur.Val < head.Val {
			cur.Next = l
			l = cur
		} else {
			r.Next = cur
			r = cur
		}
	}

	// 此时，把链表分成l...head，与head.Next...r的两段。
	// 断开两边
	head.Next = nil
	left := sortList(l)
	right := sortList(rightStart.Next)

	// head一定是第一段的结束，因此我可以直接用head连起来两个结果，因为head大于第一段的所有数字
	head.Next = right
	return left
}

// 根据 https://leetcode-cn.com/problems/sort-list/solution/sort-list-gui-bing-pai-xu-lian-biao-by-jyd/ 的解法得到
// 确实是更快的，链表排序还是用归并的好
func sortListMergeSort(head *ListNode) *ListNode {
	// 拿到长度，减少if判断
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	// 返回新的结尾，以连接本来的后面的部分
	merge := func(last *ListNode, first *ListNode, firstLen int, second *ListNode, secondLen int) *ListNode {
		for firstLen > 0 && secondLen > 0 {
			if first.Val < second.Val {
				last.Next = first
				firstLen--
				first = first.Next
			} else {
				last.Next = second
				secondLen--
				second = second.Next
			}
			last = last.Next
		}

		if firstLen > 0 {
			for firstLen > 0 {
				last.Next = first
				last = last.Next
				first = first.Next
				firstLen--
			}
		} else {
			for secondLen > 0 {
				last.Next = second
				second = second.Next
				last = last.Next
				secondLen--
			}
		}
		return last
	}

	unit := 1
	for unit < length {
		dummy := &ListNode{
			Val:  0,
			Next: head,
		}
		last := dummy
		handle := 0
		// 完整执行归并的，两段长度相等
		for handle = 0; handle+unit<<1 <= length; handle += unit << 1 {
			fast := last.Next
			slow := last.Next
			for i := 0; i < unit; i++ {
				slow = slow.Next
				fast = fast.Next.Next
			}
			// 开始归并
			firstStart := last.Next
			secondStart := slow
			firstLen := unit
			secondLen := unit
			last = merge(last, firstStart, firstLen, secondStart, secondLen)
			last.Next = fast
		}
		// 只剩下一段，无需处理。而剩下一段完整的还有另一段短的时候，则需要额外归并处理
		if length-handle > unit {
			secondLen := length - handle - unit
			secondStart := last.Next
			for i := 0; i < unit; i++ {
				secondStart = secondStart.Next
			}
			last := merge(last, last.Next, unit, secondStart, secondLen)
			last.Next = nil
		}
		head = dummy.Next
		unit = unit << 1
	}

	return head
}

func sortList(head *ListNode) *ListNode {
	return sortListMergeSort(head)
}

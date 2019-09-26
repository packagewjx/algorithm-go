package leetcode

import "container/heap"

type my23Heap []*ListNode

func (h *my23Heap) Len() int {
	return len(*h)
}

func (h *my23Heap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *my23Heap) Swap(i, j int) {
	temp := (*h)[i]
	(*h)[i] = (*h)[j]
	(*h)[j] = temp
}

func (h *my23Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *my23Heap) Pop() interface{} {
	old := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return old
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &my23Heap{}
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		h.Push(lists[i])
	}
	heap.Init(h)
	fakeHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	last := fakeHead
	for h.Len() > 0 {
		cur := heap.Pop(h).(*ListNode)
		last.Next = cur
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
		last = cur
	}
	return fakeHead.Next
}

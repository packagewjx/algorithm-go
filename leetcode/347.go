package leetcode

import "container/heap"

type my347Count struct {
	num   int
	count int
}

type my347Heap struct {
	arr []*my347Count
}

func (h *my347Heap) Len() int {
	return len(h.arr)
}

func (h *my347Heap) Less(i, j int) bool {
	return h.arr[i].count > h.arr[j].count
}

func (h *my347Heap) Swap(i, j int) {
	temp := h.arr[i]
	h.arr[i] = h.arr[j]
	h.arr[j] = temp
}

func (h *my347Heap) Push(x interface{}) {
	h.arr = append(h.arr, x.(*my347Count))
}

func (h *my347Heap) Pop() interface{} {
	temp := h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	return temp
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]*my347Count)
	for i := 0; i < len(nums); i++ {
		c, ok := count[nums[i]]
		if !ok {
			c = &my347Count{
				num:   nums[i],
				count: 0,
			}
		}
		c.count++
		count[nums[i]] = c
	}

	h := &my347Heap{arr: make([]*my347Count, 0, len(count))}
	for _, c := range count {
		h.arr = append(h.arr, c)
	}

	heap.Init(h)
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(*my347Count).num
	}
	return result
}

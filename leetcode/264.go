//+build 264heap

package leetcode

import (
	"container/heap"
)

type my264Queue []int

func (q *my264Queue) Len() int {
	return len(*q)
}

func (q *my264Queue) Less(i, j int) bool {
	return (*q)[i] < (*q)[j]
}

func (q *my264Queue) Swap(i, j int) {
	temp := (*q)[i]
	(*q)[i] = (*q)[j]
	(*q)[j] = temp
}

func (q *my264Queue) Push(x interface{}) {
	*q = append(*q, x.(int))
}

func (q *my264Queue) Pop() interface{} {
	res := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return res
}

// 维护丑数堆，每次实现仅与列表中的元素相乘，从而去除其他因子
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	// 初始化堆
	q := &my264Queue{}
	heap.Init(q)
	heap.Push(q, 1)
	last := 0
	for i := 1; i <= n; i++ {
		num := heap.Pop(q).(int)
		if num > last {
			heap.Push(q, num*2)
			heap.Push(q, num*3)
			heap.Push(q, num*5)
			last = num
		} else {
			i--
		}
	}
	return last
}

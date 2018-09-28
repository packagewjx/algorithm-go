package datastructure

import "github.com/pkg/errors"

// 一个及其简单的队列实现，用于能够预先确定队列长度的情况
// 由于是使用移动0号元素的位置进行弹出操作，因此在大量使用时，导致内存占用过多，因此不推荐
type SimpleQueue struct {
	array []interface{}
	last  int
}

func NewSimpleQueue(size int) *SimpleQueue {
	return &SimpleQueue{array: make([]interface{}, size)}
}

func (q *SimpleQueue) Poll() (interface{}, error) {
	if q.last == 0 {
		return nil, errors.New("队列中无元素")
	}

	elem := q.array[0]
	q.array = q.array[1:]
	q.last--

	return elem, nil
}

func (q *SimpleQueue) Push(item interface{}) {
	// 需要保证数组不越界
	if q.last == len(q.array) {
		q.array = append(q.array, item)
	} else {
		q.array[q.last] = item
	}
	q.last++
}

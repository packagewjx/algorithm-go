package datastructure

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewSimpleQueue(10)
	num, e := q.Poll()
	if e == nil {
		t.Errorf("此时应该出错，因为没有元素可以弹出")
	}
	q.Push(1)
	q.Push(2)
	num, _ = q.Poll()
	if num.(int) != 1 {
		t.Errorf("应该是1")
	}
	num, _ = q.Poll()
	if num.(int) != 2 {
		t.Errorf("应该是2")
	}
	_, e = q.Poll()
	if e == nil {
		t.Errorf("此时应该出错，因为没有元素可以弹出")
	}
	q.Push(3)
	q.Push(4)
	q.Push(5)
	q.Push(6)
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	q.Push(7)
	q.Push(8)
	q.Push(9)
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())

	q.Push(nil)
	fmt.Print("Push nil然后Poll:")
	fmt.Println(q.Poll())
	fmt.Println(q)
}

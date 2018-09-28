package datastructure

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntStack(t *testing.T) {
	stack := NewStack(reflect.TypeOf(0))
	var pop func() (int, error)
	var push func(int)
	stack.MakePop(&pop)
	stack.MakePush(&push)

	_, err := pop()
	if err == nil {
		t.Error("栈没有元素，应该出错")
	}
	fmt.Println(err.Error())

	// 简单测试
	push(1)
	push(2)

	i, err := pop()
	if err != nil {
		t.Error("不应出错")
	}
	if i != 2 {
		t.Error("应该为2")
	}

	push(3)
	i, err = pop()
	if err != nil {
		t.Error("不应出错")
	}
	if i != 3 {
		t.Error("应该为3")
	}

	i, err = pop()
	if err != nil {
		t.Error("不应出错")
	}
	if i != 1 {
		t.Error("应该为1")
	}

	// 测试插入多个
	for count := 0; count < 1000; count++ {
		push(count)
	}

	for stack.Len() > 0 {
		i, _ := pop()
		fmt.Println(i)
	}

	// 尝试用stack函数插入
	// TODO 无法使用pop interface
	stack.Push(1)
	iface, err := stack.Pop()
	i, ok := iface.(int)
	if !ok {
		t.Error("Pop出来的不是整数了")
	}

	for count := 0; count < 1000; count++ {
		stack.Push(count)
	}

	for stack.Len() > 0 {
		i, _ := stack.Pop()
		fmt.Println(i.(int))
	}
}

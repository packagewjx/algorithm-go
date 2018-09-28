package datastructure

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

//var (
//    empryInterfaceValue interface{}
//    emptyInterfaceType  = reflect.TypeOf(empryInterfaceValue)
//)

// 使用反射实现的栈，可以方便的构造任意类型的栈。调用MakePush和MakePop，传入u符合要求的函数指针，就可以构造两个函数
type Stack struct {
	array []reflect.Value
	typ   reflect.Type
	push  func(interface{})
	pop   func() (interface{}, error)
}

func NewStack(typ reflect.Type) *Stack {
	s := Stack{}
	s.array = make([]reflect.Value, 0, 10)
	s.MakePush(&s.push)
	s.MakePop(&s.pop)
	s.typ = typ
	return &s
}

// 构造一个Push函数，并赋值给fptr函数指针。fptr必须为这种类型的指针：func(v AnyType)
func (s *Stack) MakePush(fptr interface{}) error {
	if err := checkFuncPtr(fptr, 1, 0); err != nil {
		return err
	}

	f := reflect.ValueOf(fptr).Elem()

	push := func(v []reflect.Value) []reflect.Value {
		s.array = append(s.array, v[0])
		return nil
	}

	newFunc := reflect.MakeFunc(f.Type(), push)
	f.Set(newFunc)
	return nil
}

// 构造一个Pop函数，并赋值给fptr，fptr必须是这种类型的指针：func() (AnyType, error)
func (s *Stack) MakePop(fptr interface{}) error {
	if err := checkFuncPtr(fptr, 0, 2); err != nil {
		return err
	}

	f := reflect.ValueOf(fptr).Elem()

	pop := func(args []reflect.Value) []reflect.Value {
		var err error
		if s.Len() == 0 {
			err = errors.New("栈中无元素")
			return []reflect.Value{reflect.Zero(s.typ), reflect.ValueOf(&err).Elem()}
		}
		v := s.array[len(s.array)-1]
		s.array = s.array[:len(s.array)-1]
		//if f.Type().Out(0) == emptyInterfaceType {
		//    if !v.CanInterface() {
		//        err = errors.New("元素无法转换为Interface")
		//        return []reflect.Value{reflect.Zero(s.typ), reflect.ValueOf(&err).Elem()}
		//    }
		//    i := v.Interface()
		//    return []reflect.Value{reflect.ValueOf(&i).Elem(), reflect.ValueOf(&err).Elem()}
		//}
		return []reflect.Value{v, reflect.ValueOf(&err).Elem()}
	}

	newFunc := reflect.MakeFunc(f.Type(), pop)
	f.Set(newFunc)
	return nil
}

func (s *Stack) Push(v interface{}) error {
	if reflect.TypeOf(v) == s.typ {
		return errors.New("类型不符合")
	}
	s.push(v)
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	return s.pop()
}

func (s *Stack) Len() int {
	return len(s.array)
}

func checkFuncPtr(fptr interface{}, numIn, numOut int) error {
	value := reflect.ValueOf(fptr)
	f := value.Elem()

	// 检查是否是指针
	if value.Kind() != reflect.Ptr {
		return errors.New("传入不是函数指针")
	}

	// 检查是否是指向函数
	if f.Kind() != reflect.Func {
		return errors.New("传入不是函数指针")
	}

	// 检查输入输出数量
	if fType := f.Type(); fType.NumIn() != numIn || fType.NumOut() != numOut {
		return errors.New(fmt.Sprintf("函数输入输出数量不符合%d输入和%d输出的格式", numIn, numOut))
	}

	return nil
}

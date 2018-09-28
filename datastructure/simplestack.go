package datastructure

type SimpleStack struct {
	stack []interface{}
}

func NewSimpleStack() *SimpleStack {
	return &SimpleStack{stack: make([]interface{}, 0, 10)}
}

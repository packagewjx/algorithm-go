package leetcode

import (
	"fmt"
	"testing"
)

func TestToTreeArray(t *testing.T) {
	// 测试空数组
	treeArray, e := newTreeArray("")
	// 测试普通数组
	if e != nil {
		t.Errorf("空数组不应返回错误")
	}
	if treeArray == nil {
		t.Errorf("空数组应返回没有元素的Array")
	}
	treeArray, e = newTreeArray("[1,2]")
	if treeArray.Len() != 2 {
		t.Errorf("长度应该为2")
	}
	if val, ok := treeArray.Get(0); !ok || val != 1 {
		t.Errorf("值应该存在，且为1")
	}
	// 测试数组有误
	treeArray, e = newTreeArray("[1,2,]")
	if e == nil {
		t.Errorf("应该返回错误，由于数组出错")
	}
	// 测试只有一个元素的数组
	treeArray, e = newTreeArray("[1]")
	if treeArray.Len() != 1 {
		t.Errorf("长度应该为1")
	}
	// 测试含有null的数组
	treeArray, e = newTreeArray("[1,null,3]")
	if _, exist := treeArray.Get(1); exist {
		t.Errorf("当前值不存在，不应为true")
	}
}

func TestConstruct(t *testing.T) {
	treeArray, e := newTreeArray("[1,2,3,4,null,5]")
	if e != nil {
		t.Errorf("转换出错，信息为%s\n", e.Error())
	}
	_, err := construct(treeArray)
	if err != nil {
		t.Errorf("构造出错，信息为%s\n", e.Error())
	}
}

func TestNewTree(t *testing.T) {
	root, e := NewTree("[1,null,0,0,1]")
	if e != nil {
		t.Errorf("构造出错，信息为%s\n", e.Error())
	}
	fmt.Println(root)
}

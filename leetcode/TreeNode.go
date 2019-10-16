package leetcode

import (
	"github.com/packagewjx/algorithm-go/datastructure"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(treeString string) (*TreeNode, error) {
	newTreeArray, e := newTreeArray(treeString)
	if e != nil {
		return nil, e
	}
	return construct(newTreeArray)
}

func construct(array *treeArray) (root *TreeNode, err error) {
	val, exist := array.Get(0)
	if !exist {
		return nil, nil
	}

	root = &TreeNode{Val: val}
	// 存储构建好的node
	queue := datastructure.NewSimpleQueue(array.Len())
	queue.Push(root)

	for i := 1; i < array.Len(); {
		node, err := queue.Poll()
		if err != nil {
			return nil, errors.New("构造出错")
		}
		//if node == nil {
		//    // 如果本节点是nil，则当前两个元素应该被舍弃。注意检查越界
		//    i += 2
		//    continue
		//}
		tnode := node.(*TreeNode)

		// 左节点是一定没有超出数组长度的，不然就进不来这个循环体了
		val, exist := array.Get(i)
		if exist {
			left := &TreeNode{Val: val}
			tnode.Left = left
			queue.Push(left)
		}
		//} else {
		//    queue.Push(nil)
		//}
		i += 1
		// 判断是否有右节点
		if i < array.Len() {
			val, exist := array.Get(i)
			if exist {
				rNode := &TreeNode{Val: val}
				tnode.Right = rNode
				queue.Push(rNode)
			}
			//else {
			//    queue.Push(nil)
			//}
			i += 1
		}
	}

	return root, nil
}

// 树的数组表示
type treeArray struct {
	nums   []int
	exists []bool
}

func (ta treeArray) Len() int {
	return len(ta.nums)
}

func (ta treeArray) Get(index int) (num int, exist bool) {
	if index < 0 || index > len(ta.nums) {
		return 0, false
	}

	return ta.nums[index], ta.exists[index]
}

// 将数组[1,2,...]转换为树数组
func newTreeArray(treeString string) (*treeArray, error) {
	if len(treeString) == 0 {
		return &treeArray{nums: make([]int, 0), exists: make([]bool, 0)}, nil
	}
	treeString = strings.TrimSpace(treeString)
	// 数组表达式，检测是否符合
	matched, err := regexp.MatchString("^\\[((-?\\d+|null),)*(-?\\d+|null)]$", treeString)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.New("树数组有误")
	}

	treeString = treeString[1 : len(treeString)-1]
	numStrings := strings.Split(treeString, ",")
	nums := make([]int, len(numStrings))
	exists := make([]bool, len(numStrings))

	for i, numString := range numStrings {
		if numString == "null" {
			continue
		}

		num, err := strconv.Atoi(numString)
		if err != nil {
			return nil, err
		}

		nums[i] = num
		exists[i] = true
	}

	return &treeArray{nums, exists}, nil
}

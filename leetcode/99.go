package leetcode

import (
	"math"
)

type context struct {
	last   *TreeNode
	first  *TreeNode
	second *TreeNode
}

// 返回bool代表是否已经修复
func recoverTreeInorder(root *TreeNode, ctx *context) bool {
	// 首先处理左边
	if root.Left != nil {
		if recoverTreeInorder(root.Left, ctx) {
			return true
		}
	}

	// 本节点的处理
	if root.Val < ctx.last.Val {
		if ctx.first != nil {
			// 发现一个错误的last，如果之前已经赋值过first，说明这是第二个错误节点，此时对调
			temp := ctx.first.Val
			ctx.first.Val = root.Val
			root.Val = temp
			return true
		} else {
			// 赋值first与second
			ctx.first = ctx.last
			ctx.second = root
		}
	}

	ctx.last = root
	// 右节点处理
	if root.Right != nil {
		if recoverTreeInorder(root.Right, ctx) {
			return true
		}
	}

	return false
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	// 中序遍历，如果上一个节点比我大，则这是被对调的第一个节点。
	// 第二个节点在中序遍历中接着第一个，也可能不是
	es := &context{
		last: &TreeNode{Val: math.MinInt64},
	}
	if !recoverTreeInorder(root, es) {
		// 如果没有修复，则es的两个就是被交换的
		temp := es.first.Val
		es.first.Val = es.second.Val
		es.second.Val = temp
	}
}

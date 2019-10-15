package leetcode

import "math"

func isValidBSTInorder(root *TreeNode, last *int) bool {
	if root.Left == nil && root.Right == nil {
		// 叶子节点，查看后就返回
		if root.Val > *last {
			*last = root.Val
			return true
		} else {
			return false
		}
	}

	// 检查左节点
	if root.Left != nil && !isValidBSTInorder(root.Left, last) {
		return false
	}
	// 检查当前的
	if root.Val <= *last {
		return false
	}
	*last = root.Val
	if root.Right != nil && !isValidBSTInorder(root.Right, last) {
		return false
	}
	// 检查通过，返回true
	return true
}

// 中序遍历查看是否是二叉搜索树
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	n := math.MinInt64
	return isValidBSTInorder(root, &n)
}

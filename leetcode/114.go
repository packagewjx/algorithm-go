package leetcode

// 返回链表化的最后一个节点
func flattenBT(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}

	// 中间节点，分别链表化左边和右边，再将左边的最后连接上右边
	if root.Left != nil && root.Right != nil {
		leftLast := flattenBT(root.Left)
		rightLast := flattenBT(root.Right)
		leftLast.Left = nil
		leftLast.Right = root.Right
		root.Right = root.Left
		root.Left = nil
		return rightLast
	} else if root.Left != nil {
		// 右节点为空
		last := flattenBT(root.Left)
		root.Right = root.Left
		root.Left = nil
		return last
	} else {
		// 左节点为空
		return flattenBT(root.Right)
	}
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flattenBT(root)
}

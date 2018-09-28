package leetcode

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 判断是否是叶子节点，若是，则当是0的时候返回nil，否则返回自己
	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return nil
		} else {
			return root
		}
	}

	// 深度遍历
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	// 若自己变成了叶子，则返回是否为0
	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return nil
		} else {
			return root
		}
	} else {
		// 自己没有变成叶子，说明不是空树，返回
		return root
	}
}

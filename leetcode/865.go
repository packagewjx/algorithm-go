package leetcode

/**
后序遍历树
*/
func subtreeWithAllDeepestTraverse(root *TreeNode, curLevel int, deepest *int) (subTree *TreeNode, subTreeDeepest int) {
	if root.Left == nil && root.Right == nil {
		if curLevel > *deepest {
			*deepest = curLevel
		}
		return root, curLevel
	}

	if root.Left == nil {
		return subtreeWithAllDeepestTraverse(root.Right, curLevel+1, deepest)
	} else if root.Right == nil {
		return subtreeWithAllDeepestTraverse(root.Left, curLevel+1, deepest)
	} else {
		leftSubTree, leftSubTreeDeepest := subtreeWithAllDeepestTraverse(root.Left, curLevel+1, deepest)
		rightSubTree, rightSubTreeDeepest := subtreeWithAllDeepestTraverse(root.Right, curLevel+1, deepest)

		// 如果两边的最深都是一样的，那么返回这个root
		// 否则，返回比较深的那个subTree
		if leftSubTreeDeepest == *deepest && rightSubTreeDeepest == *deepest {
			return root, leftSubTreeDeepest
		} else if leftSubTreeDeepest < rightSubTreeDeepest {
			return rightSubTree, rightSubTreeDeepest
		} else {
			return leftSubTree, leftSubTreeDeepest
		}
	}
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	deepest := 0
	subTree, _ := subtreeWithAllDeepestTraverse(root, 0, &deepest)
	return subTree
}

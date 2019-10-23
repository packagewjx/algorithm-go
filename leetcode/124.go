package leetcode

import "math"

// 返回从root节点开始最大值的路径
func maxPathSumDFS(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := maxPathSumDFS(root.Left, max)
	right := maxPathSumDFS(root.Right, max)
	if left < 0 && right < 0 {
		// 此时，无论左边路还是右边路，还是这个为顶点的路径，加起来都肯定会更小。只看root的值即可
		if root.Val > *max {
			*max = root.Val
		}
		return root.Val
	}

	maxStart := root.Val
	if left > 0 {
		maxStart = left + root.Val
	}
	if right > 0 && right+root.Val > maxStart {
		maxStart = right + root.Val
	}
	if maxStart > *max {
		*max = maxStart
	}
	if left+right+root.Val > *max {
		*max = left + right + root.Val
	}

	return maxStart
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt64
	maxPathSumDFS(root, &max)
	return max
}

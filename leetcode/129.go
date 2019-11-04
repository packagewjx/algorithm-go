package leetcode

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0

	var dfs func(root *TreeNode, cur int)
	dfs = func(root *TreeNode, cur int) {
		cur = cur*10 + root.Val
		if root.Left == nil && root.Right == nil {
			result += cur
			return
		}
		if root.Left != nil {
			dfs(root.Left, cur)
		}
		if root.Right != nil {
			dfs(root.Right, cur)
		}
	}
	dfs(root, 0)

	return result
}

package leetcode

func preorderTree(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	preorderTree(root.Left, result)
	preorderTree(root.Right, result)
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 32)
	preorderTree(root, &result)
	return result
}

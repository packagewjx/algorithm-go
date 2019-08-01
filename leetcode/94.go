package leetcode

func inorderTraverse(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorderTraverse(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorderTraverse(root.Right, nums)
}

func inorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0, 10)
	inorderTraverse(root, &nums)
	return nums
}

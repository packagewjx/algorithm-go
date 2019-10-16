package leetcode

func pathSumDFS(root *TreeNode, sum int, cur []int, result *[][]int) {
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			res := make([]int, len(cur)+1)
			copy(res, cur)
			res[len(cur)] = root.Val
			*result = append(*result, res)
		}
		return
	}

	cur = append(cur, root.Val)
	if root.Left != nil {
		pathSumDFS(root.Left, sum-root.Val, cur, result)
	}
	if root.Right != nil {
		pathSumDFS(root.Right, sum-root.Val, cur, result)
	}
	cur = cur[:len(cur)-1]
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0, 4)
	pathSumDFS(root, sum, []int{}, &result)
	return result
}

package leetcode

func inorder230(node *TreeNode, k int, cur *int) int {
	if node == nil {
		panic("node is nil!")
	}
	if node.Left != nil {
		val := inorder230(node.Left, k, cur)
		if val != -1 {
			return val
		}
	}
	*cur++
	if k == *cur {
		return node.Val
	}
	if node.Right != nil {
		val := inorder230(node.Right, k, cur)
		if val != -1 {
			return val
		}
	}
	if *cur < k {
		return -1
	}
	panic("not possible cur cannot larger than k")
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	cur := 0
	return inorder230(root, k, &cur)
}

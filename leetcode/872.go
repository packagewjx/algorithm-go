package leetcode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	cur = 0
	return leafCheck(leafSequence(root1), root2)
}

var cur int

func leafCheck(sequence []int, root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		if root.Val == sequence[cur] {
			cur++
			return true
		} else {
			return false
		}
	}
	if root.Left != nil {
		if !leafCheck(sequence, root.Left) {
			return false
		}
	}
	if root.Right != nil {
		if !leafCheck(sequence, root.Right) {
			return false
		}
	}
	return true
}

func leafSequence(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	return append(leafSequence(root.Left), leafSequence(root.Right)...)
}

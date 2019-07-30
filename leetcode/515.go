package leetcode

func largestValuesTraverse(node *TreeNode, level int, largest map[int]int) {
	if node == nil {
		return
	}
	l, ok := largest[level]
	if !ok {
		largest[level] = node.Val
	} else if ok && l < node.Val {
		largest[level] = node.Val
	}
	largestValuesTraverse(node.Left, level+1, largest)
	largestValuesTraverse(node.Right, level+1, largest)
}

func largestValues(root *TreeNode) []int {
	largest := make(map[int]int)
	largestValuesTraverse(root, 0, largest)
	result := make([]int, len(largest))
	for i := 0; i < len(largest); i++ {
		result[i] = largest[i]
	}
	return result
}

package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	nodes := [2][]*TreeNode{
		make([]*TreeNode, 0, 16),
		make([]*TreeNode, 0, 16),
	}
	level := 0
	nodes[0] = append(nodes[0], root)

	result := make([][]int, 0, 4)
	for len(nodes[level&1]) != 0 {
		thisLevel := &nodes[level&1]
		nextLevel := &nodes[level&1^1]
		*nextLevel = (*nextLevel)[:0]

		res := make([]int, 0, len(*thisLevel))
		for i := 0; i < len(*thisLevel); i++ {
			node := (*thisLevel)[i]
			res = append(res, node.Val)
			if node.Left != nil {
				*nextLevel = append(*nextLevel, node.Left)
			}
			if node.Right != nil {
				*nextLevel = append(*nextLevel, node.Right)
			}
		}

		result = append(result, res)
		level++
	}

	return result
}

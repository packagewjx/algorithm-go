package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
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
		isLeftToRight := level&1 == 0
		*nextLevel = (*nextLevel)[:0]

		res := make([]int, 0, len(*thisLevel))
		for i := len(*thisLevel) - 1; i >= 0; i-- {
			node := (*thisLevel)[i]
			res = append(res, node.Val)
			if isLeftToRight {
				if node.Left != nil {
					*nextLevel = append(*nextLevel, node.Left)
				}
				if node.Right != nil {
					*nextLevel = append(*nextLevel, node.Right)
				}
			} else {
				// 从右往左的遍历，则反向添加子节点
				if node.Right != nil {
					*nextLevel = append(*nextLevel, node.Right)
				}
				if node.Left != nil {
					*nextLevel = append(*nextLevel, node.Left)
				}

			}
		}
		result = append(result, res)
		level++
	}

	return result
}

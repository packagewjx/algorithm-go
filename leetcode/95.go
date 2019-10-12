package leetcode

func generateBST(start, end int) []*TreeNode {
	if start == end {
		return []*TreeNode{
			{Val: start},
		}
	} else if start > end {
		return []*TreeNode{
			nil,
		}
	}

	result := make([]*TreeNode, 0, 1<<uint(end-start))
	// 依次生成i从start+1到end-1为根的树
	for i := start; i <= end; i++ {
		leftBST := generateBST(start, i-1)
		rightBST := generateBST(i+1, end)
		// 总共有len(leftBST)*len(rightBST)这么多
		for j := 0; j < len(leftBST); j++ {
			for k := 0; k < len(rightBST); k++ {
				result = append(result, &TreeNode{
					Val:   i,
					Left:  leftBST[j],
					Right: rightBST[k],
				})
			}
		}
	}
	return result
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateBST(1, n)
}

//+build 105

package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rootPos := 0
	for ; rootPos < len(inorder) && inorder[rootPos] != rootVal; rootPos++ {

	}

	root := &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:rootPos+1], inorder[:rootPos]),
		Right: buildTree(preorder[1+rootPos:], inorder[rootPos+1:]),
	}

	return root
}

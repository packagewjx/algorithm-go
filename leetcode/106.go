package leetcode

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	rootPos := 0
	for ; rootPos < len(inorder) && inorder[rootPos] != rootVal; rootPos++ {

	}

	root := &TreeNode{
		Val:   rootVal,
		Left:  buildTree(inorder[:rootPos], postorder[:rootPos]),
		Right: buildTree(inorder[rootPos+1:], postorder[rootPos:len(postorder)-1]),
	}

	return root
}

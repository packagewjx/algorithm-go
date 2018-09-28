package leetcode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}

	bigNum := 0
	bigIndex := 0

	for i, val := range nums {
		if val > bigNum {
			bigIndex = i
			bigNum = val
		}
	}

	return &TreeNode{bigNum, constructMaximumBinaryTree(nums[0:bigIndex]), constructMaximumBinaryTree(nums[bigIndex+1:])}
}

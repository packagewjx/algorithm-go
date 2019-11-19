//+build 5264

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
	root   *TreeNode
	inTree map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	// 复原
	queue := make([]*TreeNode, 0, 100)
	queue = append(queue, root)
	root.Val = 0
	inTree := make(map[int]bool)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		inTree[node.Val] = true

		if node.Left != nil {
			node.Left.Val = 2*node.Val + 1
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			node.Right.Val = 2*node.Val + 2
			queue = append(queue, node.Right)
		}
	}

	return FindElements{inTree: inTree, root: root}
}

func (this *FindElements) Find(target int) bool {
	return this.inTree[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */

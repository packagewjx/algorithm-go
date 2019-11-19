package leetcode

func postorderTraversal(root *TreeNode) []int {
	const Visited = 1
	const NotVisited = 2
	type context struct {
		node  *TreeNode
		state int
	}
	stack := make([]*context, 0, 10)
	res := make([]int, 0, 10)
	// 哨兵
	fake := &TreeNode{Val: 0, Left: nil, Right: root}
	stack = append(stack, &context{
		node:  fake,
		state: NotVisited,
	})

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.state == Visited {
			res = append(res, cur.node.Val)
		} else {
			curNode := cur.node
			// 访问我的左边，并在此过程中推栈
			for curNode != nil {
				stack = append(stack, &context{
					node:  curNode,
					state: Visited,
				})
				if curNode.Right != nil {
					stack = append(stack, &context{
						node:  curNode.Right,
						state: NotVisited,
					})
				}
				curNode = curNode.Left
			}
		}
	}
	// 多了哨兵的值，减去
	return res[:len(res)-1]
}

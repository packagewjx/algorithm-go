package leetcode

type valCoordinate struct {
	x   int
	y   int
	val int
}

func findBottomLeftValueTraverse(x, y int, node *TreeNode) *valCoordinate {
	if node.Left == nil && node.Right == nil {
		return &valCoordinate{
			x:   x,
			y:   y,
			val: node.Val,
		}
	}

	var lv, rv *valCoordinate
	if node.Left != nil {
		lv = findBottomLeftValueTraverse(x*2, y+1, node.Left)
	}
	if node.Right != nil {
		rv = findBottomLeftValueTraverse(x*2+1, y+1, node.Right)
	}
	if lv == nil {
		return rv
	}
	if rv == nil {
		return lv
	}

	if lv.y > rv.y {
		return lv
	} else if lv.y < rv.y {
		return rv
	} else {
		if lv.x < rv.x {
			return lv
		} else {
			return rv
		}
	}
}

func findBottomLeftValue(root *TreeNode) int {
	return findBottomLeftValueTraverse(1, 1, root).val
}

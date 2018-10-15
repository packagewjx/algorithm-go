package datastructure

type nodeColor bool

var RED = nodeColor(true)
var BLACK = nodeColor(false)

type rbNode struct {
	data       Data
	color      nodeColor
	parent     *rbNode
	leftChild  *rbNode
	rightChild *rbNode
}

func (n *rbNode) grandParent() *rbNode {
	if n.parent != nil {
		return n.parent.parent
	} else {
		return nil
	}
}

func (n *rbNode) uncle() *rbNode {
	gp := n.grandParent()
	if gp != nil {
		if n.parent == gp.leftChild {
			return gp.rightChild
		} else {
			return gp.leftChild
		}
	} else {
		return nil
	}
}

func (n *rbNode) key() int {
	return n.data.Key()
}

type RedBlackTree struct {
	root *rbNode
}

func (tree *RedBlackTree) Insert(data Data) {
	newNode := &rbNode{data: data, color: RED}

	// 寻找新节点的父节点，并插入
	parent, isLeft := tree.findNodeParent(newNode)
	// 插入新的节点
	if parent != nil {
		newNode.parent = parent
		if isLeft {
			parent.leftChild = newNode
		} else {
			parent.rightChild = newNode
		}
	}

	tree.changeColorToMatch(newNode)
}

func (tree *RedBlackTree) changeColorToMatch(node *rbNode) {
	// 情形1，若node为根节点
	if node.parent == nil {
		tree.root = node
		node.color = BLACK
		return
	}

	// 情形2，若父节点是黑色，则无需继续操作，直接返回
	if node.parent.color == BLACK {
		return
	}

	// 情形3，若父节点是红色，此时一定有祖父节点，因为根节点必定是黑色的。需要看情况

	// 若父节点和叔叔节点都是红色，则把他们改成黑色，且把祖父节点改成红色。
	if node.parent.color == RED && node.uncle() != nil && node.uncle().color == RED {
		node.parent.color = BLACK
		node.uncle().color = BLACK
		node.grandParent().color = RED

		// 改完后，可能祖父的父节点也是红色的，因此需要递归的修改
		tree.changeColorToMatch(node.grandParent())
		return
	}

	// 若祖父节点是黑色的，叔叔节点也是黑色的（nil也是黑色的），则需要进行旋转

	if node.grandParent().color == BLACK && (node.uncle() == nil || node.uncle().color == BLACK) {
		// - 若新节点是父节点的右节点，父节点是祖父的左节点，此时需要父节点左旋
		if node == node.parent.rightChild && node.parent == node.grandParent().leftChild {
			rbLRotation(node.parent)
			// 变换之后，node就与parent对调了，现在要处理原本的parent
			node = node.leftChild
		} else if node == node.parent.leftChild && node.parent == node.grandParent().rightChild {
			// - 若新节点是父节点的左节点，父节点是祖父的右节点，此时需要父节点右旋
			rbRRotation(node.parent)
			node = node.rightChild
		}

		if node == node.parent.leftChild && node.parent == node.grandParent().leftChild {
			// 若是左节点，且父节点也是左节点
			// 右旋祖父节点，结果是node的父节点，代替了祖父
			node = rbRRotation(node.grandParent())
			// 此时父节点就是新的最顶部节点，在原本祖父节点的位置，修改颜色即可
			node.color = BLACK
			node.rightChild.color = RED
		} else if node == node.parent.rightChild && node.parent == node.grandParent().rightChild {
			// 若是右节点，且父节点也是右节点
			node = rbLRotation(node.grandParent())
			node.color = BLACK
			node.leftChild.color = RED
		}

		// 如果已经是根节点了，则修改树的根
		if node.parent == nil {
			tree.root = node
		}
	}

}

// 旋转后返回新的根，将会设置好祖父、子节点等的关系
func rbLRotation(node *rbNode) *rbNode {
	newRoot := node.rightChild

	newRoot.parent = node.parent
	if node.parent != nil {
		if node == node.parent.rightChild {
			node.parent.rightChild = newRoot
		} else {
			node.parent.leftChild = newRoot
		}
	}

	node.rightChild = newRoot.leftChild
	if newRoot.leftChild != nil {
		node.rightChild.parent = node
	}

	node.parent = newRoot
	newRoot.leftChild = node

	return newRoot
}

func rbRRotation(node *rbNode) *rbNode {
	newRoot := node.leftChild

	newRoot.parent = node.parent
	if node.parent != nil {
		if node == node.parent.rightChild {
			node.parent.rightChild = newRoot
		} else {
			node.parent.leftChild = newRoot
		}
	}

	node.leftChild = newRoot.rightChild
	if newRoot.rightChild != nil {
		node.leftChild.parent = node
	}

	node.parent = newRoot
	newRoot.rightChild = node

	return newRoot
}

// 寻找父节点。无论key值树中是否出现，都会找到其应该插入的父节点。
// 定义红黑树的左节点的值，总小于当前节点，右节点的值大于等于当前节点的值
func (tree *RedBlackTree) findNodeParent(node *rbNode) (parent *rbNode, isLeft bool) {
	cur := tree.root

	for cur != nil {
		parent = cur
		if node.key() >= cur.key() {
			cur = cur.rightChild
			isLeft = false
		} else {
			cur = cur.leftChild
			isLeft = true
		}
	}
	return parent, isLeft
}

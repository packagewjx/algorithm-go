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

type pos byte

var LEFT = pos(1)
var RIGHT = pos(2)
var NIL = pos(3)

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

func (n *rbNode) sibling() *rbNode {
	if n.parent != nil {
		if n == n.parent.leftChild {
			return n.rightChild
		} else {
			return n.leftChild
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

// 插入数据
func (tree *RedBlackTree) Insert(data Data) {
	newNode := &rbNode{data: data, color: RED}

	// 寻找新节点的父节点，并插入
	parent, isLeft := findNodeParent(newNode.key(), tree.root)
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

// 删除并返回键值为key的数据
func (tree *RedBlackTree) Delete(key int) Data {
	deleteNode := tree.getNode(key)
	data := deleteNode.data

	// 若是双节点子女，则需要先找到最小值，替换
	if deleteNode.leftChild != nil && deleteNode.rightChild != nil {
		// 找到deleteNode右子树中的最小值
		cur := deleteNode.rightChild
		for cur.leftChild != nil {
			cur = cur.leftChild
		}
		// 用cur的替换当前的deleteNode
		deleteNode.data = cur.data
		// 要删除的节点变成了这个最小值
		deleteNode = cur
	}

	// 现在转换成删除只有一个子女的中间节点或叶子的问题
	tree.deleteOneChildNode(deleteNode)

	return data
}

func (tree *RedBlackTree) Get(key int) Data {
	return tree.getNode(key).data
}

func (tree *RedBlackTree) getNode(key int) *rbNode {
	cur := tree.root
	for cur != nil && cur.key() != key {
		if cur.key() > key {
			cur = cur.leftChild
		} else /*key >= cur.key*/ {
			cur = cur.rightChild
		}
	}

	if cur == nil {
		return nil
	} else {
		return cur
	}
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

// 删除只有一个儿子或叶子节点
func (tree *RedBlackTree) deleteOneChildNode(node *rbNode) *rbNode {
	// 检查是不是只有一个或空
	if node.leftChild != nil && node.rightChild != nil {
		panic("此函数不能删除双儿子")
	}

	deleteNode := node

	var child *rbNode

	// node在父节点中的位置
	var parentPos pos

	// 给child、parent、sibling赋值
	if node.leftChild != nil {
		child = node.leftChild
	} else if node.rightChild != nil {
		child = node.rightChild
	} else {
		child = nil
	}

	if node.parent != nil {
		if node == node.parent.leftChild {
			parentPos = LEFT
		} else {
			parentPos = RIGHT
		}
	} else {
		parentPos = NIL
	}

	// 若删除红色的节点，直接用儿子代替即可
	if node.color == RED {
		child.parent = node.parent
		if parentPos == LEFT {
			node.parent.leftChild = child
		} else if parentPos == RIGHT {
			node.parent.rightChild = child
		}
		return deleteNode
	}

	// 若节点是黑色，而儿子是红色
	if node.color == BLACK && child != nil && child.color == RED {
		child.color = BLACK

		child.parent = node.parent
		if parentPos == LEFT {
			node.parent.leftChild = child
		} else if parentPos == RIGHT {
			node.parent.rightChild = child
		}
		return deleteNode
	}

	// 接下来的情况是，节点是黑色的，儿子也是黑色的
	// 首先做的是，将节点替换成儿子
	var nilNode *rbNode
	if child == nil {
		// 使用nilNode处理nil情况
		nilNode = &rbNode{color: BLACK}
		child = nilNode
	}
	child.parent = node.parent

	if parentPos == LEFT {
		node.parent.leftChild = child
	} else {
		node.parent.rightChild = child
	}

	// 进入递归函数处理
	deleteNode = tree.deleteBlackNodeAndChild(child)
	if nilNode != nil {
		// 若使用了nilNode处理
	}
	// TODO 没有处理空节点的情况
	return nil
}

// 删除黑儿子和自己黑的时候。注意这里实际上已经被删除，node是本来的儿子
func (tree *RedBlackTree) deleteBlackNodeAndChild(node *rbNode) *rbNode {
	var parentPos pos

	if node.parent != nil {
		if node == node.parent.leftChild {
			parentPos = LEFT
		} else {
			parentPos = RIGHT
		}
	} else {
		return node
	}

	// 如果兄弟节点是红色的，则将父节点旋转旋，使得兄弟节点取代父节点位置
	if node.sibling() != nil && node.sibling().color == RED {
		if parentPos == LEFT {
			rbLRotation(node.parent)
		} else {
			rbRRotation(node.parent)
		}
		// 继续处理
		node.parent.color = RED
		node.sibling().color = BLACK
	}

	if node.parent.color == BLACK && node.sibling() != nil && node.sibling().color == BLACK &&
		(node.sibling().leftChild == nil || node.sibling().leftChild.color == BLACK) &&
		(node.sibling().rightChild == nil || node.sibling().rightChild.color == BLACK) {
		// 如果兄弟节点，及其子女都是黑色，则需要把兄弟改成红色节点，父亲也是黑色
		// 并在父亲上继续本删除
		node.sibling().color = RED
		return tree.deleteBlackNodeAndChild(node.parent)
	}

	// 若兄弟节点及其儿子都是黑色，而父亲是红色
	// 直接替换兄弟与父亲的颜色即可
	if node.parent.color == RED && node.sibling() != nil && node.sibling().color == BLACK &&
		(node.sibling().leftChild == nil || node.sibling().leftChild.color == BLACK) &&
		(node.sibling().rightChild == nil || node.sibling().rightChild.color == BLACK) {
		node.parent.color = BLACK
		node.sibling().color = RED
	}

	// 兄弟节点是黑色，但是其中一个儿子是红色
	if node.sibling() != nil && node.sibling().color == BLACK {
		// 首先改变颜色
		node.sibling().color = RED
		if node.sibling().leftChild != nil && node.sibling().leftChild.color == RED {
			// 左儿子是红色，则右旋，让左儿子替代兄弟位置
			rbRRotation(node.sibling())
		} else if node.sibling().rightChild != nil && node.sibling().rightChild.color == RED {
			rbLRotation(node.sibling())
		}
		// 兄弟节点变了，改成黑色
		node.sibling().color = BLACK
	}

	// 如果兄弟节点中，有红色子节点，且兄弟是右子节点且其红色子节点也是右子节点或相反
	if node.sibling() != nil && node.sibling().color == BLACK &&
		node.sibling() == node.parent.rightChild && node.sibling().rightChild != nil &&
		node.sibling().rightChild.color == RED {
		rbLRotation(node.parent)
		// 交换颜色
		node.grandParent().color = node.parent.color
		node.parent.color = BLACK
	} else if node.sibling() != nil && node.sibling().color == BLACK &&
		node.sibling() == node.parent.leftChild && node.sibling().leftChild != nil &&
		node.sibling().leftChild.color == RED {
		rbRRotation(node.parent)
		node.grandParent().color = node.parent.color
		node.parent.color = BLACK
	}

	return nil
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
func findNodeParent(key int, root *rbNode) (parent *rbNode, isLeft bool) {
	cur := root

	for cur != nil {
		parent = cur
		if key >= cur.key() {
			cur = cur.rightChild
			isLeft = false
		} else {
			cur = cur.leftChild
			isLeft = true
		}
	}
	return parent, isLeft
}

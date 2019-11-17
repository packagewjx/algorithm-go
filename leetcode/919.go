//+build 919

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	curLevel       int
	root           *TreeNode
	last2LevelNode []*TreeNode
	lastLevelNode  []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	if root == nil {
		return CBTInserter{
			curLevel:       0,
			root:           nil,
			last2LevelNode: make([]*TreeNode, 0, 16),
			lastLevelNode:  make([]*TreeNode, 0, 16),
		}
	}

	inserter := CBTInserter{
		curLevel:       0,
		root:           nil,
		last2LevelNode: nil,
		lastLevelNode:  nil,
	}

	nodes := make([][]*TreeNode, 2)
	nodes[0] = make([]*TreeNode, 1, 16)
	nodes[0][0] = root
	nodes[1] = make([]*TreeNode, 0, 16)
	NextLevelNodesIndex := 1
	curLevel := 0
	// NextLevelNodesIndex^1就是CurrentLevelNodesIndex
	for len(nodes[NextLevelNodesIndex^1]) > 0 {
		if nodes[NextLevelNodesIndex^1][0].Left == nil {
			// 说明没有下一层了可以退出。此时nodes[NextLevelNodesIndex^1]就是最后一层，而nodes[NextLevelNodesIndex]则是上层
			break
		}
		nodes[NextLevelNodesIndex] = nodes[NextLevelNodesIndex][:0]
		for i := 0; i < len(nodes[NextLevelNodesIndex^1]); i++ {
			node := nodes[NextLevelNodesIndex^1][i]
			if node.Right != nil {
				nodes[NextLevelNodesIndex] = append(nodes[NextLevelNodesIndex], node.Left, node.Right)
			} else if node.Left != nil {
				nodes[NextLevelNodesIndex] = append(nodes[NextLevelNodesIndex], node.Left)
				// 下一个就没有子女了，可以退出
				break
			} else {
				// 空的也退出
				break
			}
		}
		NextLevelNodesIndex = NextLevelNodesIndex ^ 1
		curLevel++
	}

	inserter.last2LevelNode = nodes[NextLevelNodesIndex]
	inserter.lastLevelNode = nodes[NextLevelNodesIndex^1]
	// cur代表当前需要插入的完全二叉树数组的位置
	inserter.curLevel = curLevel
	inserter.root = root
	return inserter
}

func (this *CBTInserter) Insert(v int) int {
	// 处理root是空的情况
	if this.root == nil {
		this.root = &TreeNode{
			Val:   v,
			Left:  nil,
			Right: nil,
		}
		this.curLevel = 1
		this.last2LevelNode = append(this.last2LevelNode, this.root)
		return v
	}

	node := &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}

	// 本代码段判断是否是新的一层
	curLevelFull := 1 << uint(this.curLevel)
	if curLevelFull == len(this.lastLevelNode) {
		// 说明这一层满了。下面交换两个数组
		temp := this.last2LevelNode
		this.last2LevelNode = this.lastLevelNode
		this.lastLevelNode = temp[:1]
		this.last2LevelNode[0].Left = node
		this.lastLevelNode[0] = node
		this.curLevel++
		return this.last2LevelNode[0].Val
	}
	// 这一层没满
	index := len(this.lastLevelNode)
	this.lastLevelNode = append(this.lastLevelNode, node)
	if index&1 == 1 {
		this.last2LevelNode[index/2].Right = node
	} else {
		this.last2LevelNode[index/2].Left = node
	}
	return this.last2LevelNode[index/2].Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

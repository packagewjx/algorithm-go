package leetcode

var zeroNodeNext = make([]*GraphNode, 0, 0)

type GraphNode struct {
	Next []*GraphNode
	Val  int
}

func NewGraph(graph [][]int) map[int]*GraphNode {
	graphNodeMap := make(map[int]*GraphNode)
	for i, nodeNext := range graph {
		gNode, exist := graphNodeMap[i]
		if !exist {
			gNode = &GraphNode{Val: i, Next: make([]*GraphNode, len(nodeNext))}
			graphNodeMap[i] = gNode
		} else {
			gNode.Next = make([]*GraphNode, len(nodeNext))
		}

		if len(nodeNext) == 0 {
			gNode.Next = zeroNodeNext
		}

		for i, val := range nodeNext {
			node, exist := graphNodeMap[val]
			if !exist {
				node = &GraphNode{Val: val, Next: nil}
				graphNodeMap[val] = node
			}
			gNode.Next[i] = node
		}
	}

	return graphNodeMap
}

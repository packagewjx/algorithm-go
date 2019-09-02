package leetcode

// 本地能通过，leetcode却错误

type my684Edge struct {
	end1     *my684Node
	end2     *my684Node
	priority int
}

type my684Node struct {
	val  int
	next []*my684Edge
}

func findCyclicPath(cur *my684Node, path []*my684Edge, pathNode []int, lastNode *my684Node) (bool, []*my684Edge) {
	curNodeIndex := len(path)
	pathNode[cur.val] = curNodeIndex
	for i := 0; i < len(cur.next); i++ {
		edge := cur.next[i]
		var next *my684Node
		if edge.end1 == cur {
			next = edge.end2
		} else {
			next = edge.end1
		}
		if next == lastNode {
			continue
		}
		// 查看next是否已经出现，若是，则找到了回路
		if pathNode[next.val] != -1 {
			path = append(path, edge)
			path = path[pathNode[next.val]:]
			return true, path
		}
		path = append(path, edge)
		found, newPath := findCyclicPath(next, path, pathNode, cur)
		if found {
			return true, newPath
		} else {
			//需要切断
			path = path[:curNodeIndex]
		}
	}
	pathNode[cur.val] = -1
	return false, path
}

func findRedundantConnection(edges [][]int) []int {
	nodes := make(map[int]*my684Node)
	for i := 0; i < len(edges); i++ {
		fromNode, ok := nodes[edges[i][0]]
		if !ok {
			fromNode = &my684Node{
				val:  edges[i][0],
				next: make([]*my684Edge, 0, 10),
			}
			nodes[edges[i][0]] = fromNode
		}
		toNode, ok := nodes[edges[i][1]]
		if !ok {
			toNode = &my684Node{
				val:  edges[i][1],
				next: make([]*my684Edge, 0, 10),
			}
			nodes[edges[i][1]] = toNode
		}
		edge := &my684Edge{
			end1:     fromNode,
			end2:     toNode,
			priority: i,
		}
		fromNode.next = append(fromNode.next, edge)
		toNode.next = append(toNode.next, edge)
	}

	var foundPath []*my684Edge
	for i := 0; i < len(nodes); i++ {
		pathNode := make([]int, len(nodes)+1)
		for i := 0; i < len(pathNode); i++ {
			pathNode[i] = -1
		}
		path := make([]*my684Edge, 0, len(nodes))
		cur := nodes[i+1]
		found, cyclic := findCyclicPath(cur, path, pathNode, nil)
		if found {
			foundPath = cyclic
			break
		}
	}

	if foundPath == nil {
		panic("找不到回路！")
	}

	// 此时的path是回路，返回最大的priority的
	biggest := foundPath[0]
	for i := 1; i < len(foundPath); i++ {
		if foundPath[i].priority > biggest.priority {
			biggest = foundPath[i]
		}
	}
	return []int{biggest.end1.val, biggest.end2.val}
}

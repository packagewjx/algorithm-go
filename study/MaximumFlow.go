package study

import . "github.com/packagewjx/algorithm-go/datastructure"

type flowVertex struct {
	*WeightedDirectedVertex
	// 后向边及其权重
	Next map[*flowVertex]int
	Back map[*flowVertex]int

	// 存储已用流量，正数代表前向边已用，负数为后向边已用
	flowUsed map[*flowVertex]int

	// 标记，分别是从哪个来，以及是前向流量还是后向流量
	flowFrom *flowVertex
	// 这次遍历从flowFrom过来的流量大小
	currentFlow int
}

func MaximumFlow(vertices map[int]*WeightedDirectedVertex, source *WeightedDirectedVertex, dest *WeightedDirectedVertex) int {
	// 计算后向边
	flowVertices := toFlowVertices(vertices)
	fSource := flowVertices[source.Key()]
	fDest := flowVertices[dest.Key()]

	// 标记源点
	fSource.currentFlow = 0x7FFFFFFF
	fSource.flowFrom = fSource

	queue := make([]*flowVertex, 1, 10)
	queue[0] = fSource

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		// 前向边遍历
		for next, forwardFlow := range vertex.Next {
			if next.flowFrom == nil {
				// 这代表没有被标记
				// 计算本次的flow
				thisFlow := forwardFlow - vertex.flowUsed[next]
				if thisFlow > 0 {
					// 需要用绝对值判断
					absVertexFlow := vertex.currentFlow
					if absVertexFlow < 0 {
						absVertexFlow = -absVertexFlow
					}

					// 标记next可以从vertex来，流量为thisFlow或本节点currentFlow的最小值
					if thisFlow < absVertexFlow {
						next.currentFlow = thisFlow
					} else {
						next.currentFlow = absVertexFlow
					}
					next.flowFrom = vertex
					queue = append(queue, next)
				}
			}
		}

		// 后向边遍历
		for back := range vertex.Back {
			if back.flowFrom == nil {
				// 若这个点还没有被标记
				// 得到前向边（back到vertex）来的流量
				thisFlow := back.flowUsed[vertex]
				if thisFlow > 0 {
					abs := vertex.currentFlow
					if abs < 0 {
						abs = -abs
					}

					// 小于0代表后向流量大于0
					// 最后需要在前向边来的流量减去currentFlow
					if thisFlow < abs {
						back.currentFlow = -thisFlow
					} else {
						back.currentFlow = -abs
					}
					back.flowFrom = vertex
					queue = append(queue, back)
				}
			}
		}

		if fDest.flowFrom != nil {
			// 汇点被标记
			cur := fDest
			for cur != fSource {
				if cur.currentFlow > 0 {
					cur.flowFrom.flowUsed[cur] = cur.flowFrom.flowUsed[cur] + cur.currentFlow
				} else {
					cur.flowUsed[cur.flowFrom] = cur.flowUsed[cur.flowFrom] + cur.currentFlow
				}
				cur = cur.flowFrom
			}

			// 除去标记
			for _, vertex := range flowVertices {
				if vertex == fSource {
					continue
				}

				vertex.currentFlow = 0
				vertex.flowFrom = nil
			}

			// 重新初始化Queue
			queue = queue[:1]
			queue[0] = fSource
		}
	}

	sum := 0
	for vertex := range fDest.Back {
		sum += vertex.flowUsed[fDest]
	}
	return sum
}

func toFlowVertices(vertices map[int]*WeightedDirectedVertex) map[int]*flowVertex {
	result := make(map[int]*flowVertex, len(vertices))

	// 先创建
	for key, currentVertex := range vertices {
		result[key] = &flowVertex{WeightedDirectedVertex: currentVertex, Back: map[*flowVertex]int{}, Next: map[*flowVertex]int{}, flowUsed: map[*flowVertex]int{}}
	}

	// 然后设置Back和Next
	for _, currentVertex := range result {
		originalNext := currentVertex.WeightedDirectedVertex.Next
		for nextVertex, weight := range originalNext {
			nextVertexNow := result[nextVertex.Key()]
			currentVertex.Next[nextVertexNow] = weight
			nextVertexNow.Back[currentVertex] = weight
		}
	}

	return result
}

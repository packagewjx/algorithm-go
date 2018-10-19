package study

import . "github.com/packagewjx/algorithm-go/datastructure"

// 找出最大二分图匹配
func MaximumBipartiteMatching(graph *MatrixBasedGraph) map[int]int {
	//matching, matched := findSomeMatching(graph)
	return nil
}

// 找出这个二分图其中一个集合
func findBipartiteVertices(graph *MatrixBasedGraph) ([]int, []int) {
	// 分割集合，键为v，值为集合的编号，若是二分图，最终应该只有两个集合编号
	U := map[int]bool{}
	V := map[int]bool{}

	for v := 0; v < graph.Len(); v++ {
		if U[v] || V[v] {
			// 若已经在某个集合，则跳过这个点
			continue
		}
		var adjSet *map[int]bool
		for adv := range graph.Adjacent(v) {
			// v与其邻接点，应该不属于同一个集合
			if adjSet != nil {
				// 检查邻接点的集合，应该都是同一个
				if (adjSet == &U && V[adv]) || (adjSet == &V && U[adv]) {
					panic("邻接点不是同一个集合的！")
				}
				// 检查这个点是否已经加入了集合，然后计入
				_, isU := U[adv]
				_, isV := V[adv]
				if !isU && !isV {
					(*adjSet)[adv] = true
				}
			} else {
				// 邻接点集合未确定，确定邻接点的集合
				if U[adv] {
					adjSet = &U
				} else if V[adv] {
					adjSet = &V
				} else {
					// 这里v和adv都没有被加入集合
					U[v] = true
					V[adv] = true
					adjSet = &V
				}
			}
		}
	}

	uVertices := make([]int, 0, len(U))
	for u := range U {
		uVertices = append(uVertices, u)
	}

	vVertices := make([]int, 0, len(V))
	for v := range V {
		vVertices = append(vVertices, v)
	}

	return uVertices, vVertices
}

func findSomeMatching(graph *MatrixBasedGraph) (matching map[int]int, matched map[int]bool) {
	matched = map[int]bool{}
	matching = map[int]int{}

	for v := 0; v < graph.Len(); v++ {
		if matched[v] {
			continue
		}

		adjacent := graph.Adjacent(v)
		for adv, _ := range adjacent {
			if matched[adv] {
				continue
			}
			matching[v] = adv
			matched[v] = true
			matched[adv] = true
		}
	}

	return matching, matched
}

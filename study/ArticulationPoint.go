package study

import . "github.com/packagewjx/algorithm-go/datastructure"

var visited map[*Vertex]int
var low map[*Vertex]int
var visitNum int
var articulationPoints map[*Vertex]bool

func ArticulationPoint(v *Vertex) []*Vertex {
	articulationPoints = map[*Vertex]bool{}
	visited = map[*Vertex]int{}
	low = map[*Vertex]int{}
	visitNum = 0

	dfs(v, nil)

	ap := []*Vertex{}
	for key, _ := range articulationPoints {
		ap = append(ap, key)
	}
	return ap
}

func dfs(v *Vertex, predecessor *Vertex) {
	visited[v] = visitNum
	visitNum++
	low[v] = visited[predecessor]

	for _, adj := range v.Adjacent {
		if adjVisited, ok := visited[adj]; ok {
			if adj == predecessor {
				// 若是从这个点过来的就不用再执行了，已经赋过值
				continue
			}
			// 若已经访问过，则看这个点的visited是不是比low[v]要小
			if adjVisited < low[v] {
				// 是的话，说明v可以访问这个visited更小的点，更新low[v]
				low[v] = adjVisited
			}
			continue
		}

		// 若没有访问，则继续访问下去
		dfs(adj, v)

		if low[adj] >= visited[v] {
			// 如果我的子节点能够访问的最低visited节点大于等于我，就是说明子节点最多能访问
			// 到我，因此我是关节点
			articulationPoints[v] = true
		} else if low[adj] < low[v] {
			// 如果我的子节点能够去比我更低的地方，说明我也可以去这个更低的地方，
			// 因此更新low[v]
			low[v] = low[adj]
		}
	}

	if predecessor == nil && len(v.Adjacent) > 1 {
		articulationPoints[v] = true
	}

}

package study

import (
	"container/heap"
	. "github.com/packagewjx/algorithm-go/datastructure"
)

type edgeMinHeap struct {
	Array [][]int
	Graph *MatrixBasedGraph
}

func (heap *edgeMinHeap) Push(x interface{}) {
	heap.Array = append(heap.Array, x.([]int))
}

func (heap *edgeMinHeap) Pop() interface{} {
	temp := heap.Array[heap.Len()-1]
	heap.Array = heap.Array[:heap.Len()-1]
	return temp
}

func (heap *edgeMinHeap) Len() int {
	return len(heap.Array)
}

func (heap *edgeMinHeap) Less(i, j int) bool {
	return heap.Graph.Weight(heap.Array[i][0], heap.Array[i][1]) < heap.Graph.Weight(heap.Array[j][0], heap.Array[j][1])
}

func (heap *edgeMinHeap) Swap(i, j int) {
	temp := heap.Array[i]
	heap.Array[i] = heap.Array[j]
	heap.Array[j] = temp
}

func Prim(graph *MatrixBasedGraph) *EdgeSet {
	vertices := NewVertexSet()
	edges := NewEdgeSet()

	// 使用最小堆存储所有边
	minHeap := &edgeMinHeap{Graph: graph, Array: [][]int{}}

	heap.Init(minHeap)
	// 加入0点
	vertices.Add(0)
	// 加入0点的所有边
	addAllEdgeToHeap(minHeap, graph, 0)

	for i := 1; i < graph.Len(); i++ {
		smallestEdge := heap.Pop(minHeap).([]int)
		for vertices.Exist(smallestEdge[0]) && vertices.Exist(smallestEdge[1]) {
			// 边上两个顶点不能都在点集内，因此一直pop找到不在的
			smallestEdge = heap.Pop(minHeap).([]int)
		}

		// 找到后，一定是从点集内的点出发，到达点集外的，因此加入1的点
		vertices.Add(smallestEdge[1])
		edges.Add(smallestEdge[0], smallestEdge[1])
		addAllEdgeToHeap(minHeap, graph, smallestEdge[1])
	}

	return edges
}

// 加入v的所有边到minHeap中
func addAllEdgeToHeap(minHeap *edgeMinHeap, graph *MatrixBasedGraph, v int) {
	for adv := range graph.Adjacent(v) {
		heap.Push(minHeap, []int{v, adv})
	}
}

func Kruscal(graph *MatrixBasedGraph) *EdgeSet {
	// 连通性
	connectivity := make([][]bool, graph.Len())
	minHeap := &edgeMinHeap{Graph: graph, Array: make([][]int, 0, 10)}
	heap.Init(minHeap)

	for i := 0; i < graph.Len(); i++ {
		connectivity[i] = make([]bool, graph.Len())
		// 对角线赋值为true
		connectivity[i][i] = true

		for j := i + 1; j < graph.Len(); j++ {
			if graph.Weight(i, j) > 0 {
				heap.Push(minHeap, []int{i, j})
			}
		}
	}

	edgeSet := NewEdgeSet()

	encounter := 0
	for encounter < graph.Len()-1 {
		edge := heap.Pop(minHeap).([]int)
		if noCycle(edge, connectivity) {
			// 加入edge到路径集合
			edgeSet.Add(edge[0], edge[1])

			// 更新连通性
			updateConnectivity(edge, connectivity)
			encounter++
		}
	}

	return edgeSet
}

func updateConnectivity(edge []int, connectivity [][]bool) {
	// 两个点能够互相连通了
	connectivity[edge[0]][edge[1]] = true
	connectivity[edge[1]][edge[0]] = true

	// edge[0]能连通的所有点，edge[1]都可以去了
	for i := 0; i < len(connectivity[edge[0]]); i++ {
		// 或运算在这里保证了，本来连通的两个点，不被connectivity[edge[0]][i]的不连通性更新
		// 换句话说，要么是本来已经连通，要么是edge[0]所带来了新的连通性
		connectivity[i][edge[1]] = connectivity[i][edge[1]] || connectivity[edge[0]][i]
		connectivity[edge[1]][i] = connectivity[i][edge[1]]

		connectivity[edge[0]][i] = connectivity[edge[0]][i] || connectivity[edge[1]][i]
		connectivity[i][edge[0]] = connectivity[edge[0]][i]
	}
}

func noCycle(edge []int, connectivity [][]bool) bool {
	// 两个顶点没有连通，则是没有回路产生
	return !connectivity[edge[0]][edge[1]]
}

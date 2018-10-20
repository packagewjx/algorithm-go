package datastructure

import (
	"fmt"
)

type Vertex struct {
	Data     Data
	Adjacent []*Vertex
}

func (v *Vertex) String() string {
	Adjacent := ""
	for _, value := range v.Adjacent {
		Adjacent += fmt.Sprintf("%d,", value.Data.Key())
	}
	Adjacent = Adjacent[:len(Adjacent)-1]
	return fmt.Sprintf("%d:[%s]", v.Data.Key(), Adjacent)
}

func BuildVerticesFromMap(adjacentMap map[int][]int) map[int]*Vertex {
	vertices := make(map[int]*Vertex)

	for key, _ := range adjacentMap {
		vertices[key] = &Vertex{Data: &Dummy{K: key}, Adjacent: []*Vertex{}}
	}

	for key, v := range vertices {
		adjacent := adjacentMap[key]
		for _, Adjacent := range adjacent {
			v.Adjacent = append(v.Adjacent, vertices[Adjacent])
		}
	}

	return vertices
}

// PROBLEM
func BuildVerticesFromMatrix(adjacentMatrix [][]bool) map[int]*Vertex {
	vertices := make(map[int]*Vertex)

	getOrCreate := func(key int) (v *Vertex) {
		if mv, ok := vertices[key]; !ok {
			v = &Vertex{Data: &Dummy{K: key}, Adjacent: []*Vertex{}}
			vertices[key] = v
		} else {
			v = mv
		}
		return
	}

	for cur, adjacent := range adjacentMatrix {
		// 先从已有点图查看是否之前已经创建了
		v := getOrCreate(cur)

		for _, isAdjacent := range adjacent {
			if isAdjacent {
				av := getOrCreate(cur)
				v.Adjacent = append(v.Adjacent, av)
			}
		}
	}

	return vertices
}

type WeightedDirectedVertex struct {
	Data Data
	// 由本顶点指向的下一个顶点，和权重
	Next map[*WeightedDirectedVertex]int
}

func (w *WeightedDirectedVertex) Key() int {
	return w.Data.Key()
}

func BuildFromMatrix(matrix [][]int) map[int]*WeightedDirectedVertex {
	vertices := make(map[int]*WeightedDirectedVertex, len(matrix))

	// 先创建所有顶点
	for i := 0; i < len(matrix); i++ {
		vertices[i] = &WeightedDirectedVertex{Data: &Dummy{K: i}, Next: map[*WeightedDirectedVertex]int{}}
	}

	for i := 0; i < len(matrix); i++ {
		cur := vertices[i]
		for j := 0; j < len(matrix[i]); j++ {
			if i == j {
				continue
			}

			if matrix[i][j] > 0 {
				next := vertices[j]
				cur.Next[next] = matrix[i][j]
			}
		}
	}

	return vertices
}

type MatrixBasedGraph struct {
	// 第一个下标为起始点，第二个为终点，值为权重
	Matrix [][]int
}

func NewMatrixBasedGraphUsingList(list map[int][]int) *MatrixBasedGraph {
	graph := &MatrixBasedGraph{Matrix: make([][]int, len(list), len(list))}
	for v, adj := range list {
		graph.Matrix[v] = make([]int, len(list), len(list))

		for _, adv := range adj {
			graph.Matrix[v][adv] = 1
		}
	}

	return graph
}

// 返回图，key为vertex的序号，值为边的权重
func (g *MatrixBasedGraph) Adjacent(v int) map[int]int {
	adj := make(map[int]int, 10)
	for key, weight := range g.Matrix[v] {
		if weight > 0 {
			adj[key] = weight
		}
	}
	return adj
}

func (g *MatrixBasedGraph) Weight(start, end int) int {
	return g.Matrix[start][end]
}

func (g *MatrixBasedGraph) Len() int {
	return len(g.Matrix)
}

// 有向图中求反向边的函数，找出哪些顶点有到v的路径
func (g *MatrixBasedGraph) BackwardEdge(v int) map[int]int {
	back := make(map[int]int, 10)
	for i := 0; i < len(g.Matrix); i++ {
		if g.Matrix[i][v] > 0 {
			back[i] = g.Matrix[i][v]
		}
	}
	return back
}

type VertexSet struct {
	Vertices map[int]bool
}

func NewVertexSet() *VertexSet {
	return &VertexSet{Vertices: map[int]bool{}}
}

func (set *VertexSet) Add(v int) {
	set.Vertices[v] = true
}

// 查看某个顶点v是否在点集中
func (set *VertexSet) Exist(v int) bool {
	return set.Vertices[v]
}

func (set *VertexSet) AllPoint() []int {
	result := make([]int, 0, len(set.Vertices))

	for v, exist := range set.Vertices {
		if exist {
			result = append(result, v)
		}
	}

	return result
}

type EdgeSet struct {
	// 边以长度为2数组形式保存
	Edges map[int][]int
}

func NewEdgeSet() *EdgeSet {
	return &EdgeSet{map[int][]int{}}
}

func (set *EdgeSet) Add(start, end int) {
	hash := start ^ end
	for edge, ok := set.Edges[hash]; ok && edge[0] != start && edge[1] != end; edge, ok = set.Edges[hash] {
		// 解决冲突
		hash = hash + start + 1
	}

	// 保存数据
	set.Edges[hash] = []int{start, end}
}

func (set *EdgeSet) Exist(start, end int) bool {
	hash := start ^ end
	var ok bool
	var edge []int
	for edge, ok = set.Edges[hash]; ok && edge[0] != start && edge[1] != end; edge, ok = set.Edges[hash] {
		// 往下查找
		hash = hash + start + 1
	}

	if !ok {
		return false
	} else if edge[0] == start && edge[1] == end {
		return true
	}
	return false
}

func (set *EdgeSet) AllEdges() [][]int {
	result := make([][]int, 0, len(set.Edges))
	for _, edge := range set.Edges {
		result = append(result, edge)
	}
	return result
}

func (set *EdgeSet) Len() int {
	return len(set.Edges)
}

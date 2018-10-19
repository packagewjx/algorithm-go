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

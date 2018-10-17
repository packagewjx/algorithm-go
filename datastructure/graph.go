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

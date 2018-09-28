package study

import (
	"fmt"
	"strconv"
)

type vertex struct {
	name     string
	adjacent []*vertex
}

func (v *vertex) String() string {
	adjacents := ""
	for _, value := range v.adjacent {
		adjacents += value.name + " "
	}
	adjacents = adjacents[:len(adjacents)-1]
	return fmt.Sprintf("%s:[%s]", v.name, adjacents)
}

func buildFromMap(adjacentMap map[string][]string) map[string]*vertex {
	vertices := make(map[string]*vertex)

	for name, _ := range adjacentMap {
		vertices[name] = &vertex{name: name, adjacent: []*vertex{}}
	}

	for name, v := range vertices {
		adjacents := adjacentMap[name]
		for _, adjacent := range adjacents {
			v.adjacent = append(v.adjacent, vertices[adjacent])
		}
	}

	return vertices
}

func buildFromMatrix(adjacentMatrix [][]bool) map[string]*vertex {
	vertices := make(map[string]*vertex)

	getOrCreate := func(name string) (v *vertex) {
		if mv, ok := vertices[name]; !ok {
			v = &vertex{name: name, adjacent: []*vertex{}}
			vertices[name] = v
		} else {
			v = mv
		}
		return
	}

	for cur, adjacents := range adjacentMatrix {
		name := strconv.Itoa(cur)
		// 先从已有点图查看是否之前已经创建了
		v := getOrCreate(name)

		for vName, isAdjacent := range adjacents {
			if isAdjacent {
				av := getOrCreate(strconv.Itoa(vName))
				v.adjacent = append(v.adjacent, av)
			}
		}
	}

	return vertices
}

var visited map[*vertex]bool

func simpleDFS(v *vertex) {
	visited = make(map[*vertex]bool)
	simpleDFSRecursive(v)
}

func simpleDFSRecursive(v *vertex) {
	fmt.Println(v.name)
	visited[v] = true
	for _, val := range v.adjacent {
		if !visited[val] {
			simpleDFSRecursive(val)
		}
	}
}

var predecessor map[*vertex]*vertex
var discovery map[*vertex]int

// 在DFS树中，某个点v经过其自己或者后代（discovery比v大）能够到达的最先被访问过的点的discovery数字
var lowerDiscovery map[*vertex]int
var articulationPoints []*vertex
var time int

func articulationPoint(v *vertex) []*vertex {
	predecessor = map[*vertex]*vertex{}
	discovery = map[*vertex]int{}
	time = 1
	articulationPoints = []*vertex{}
	lowerDiscovery = map[*vertex]int{}

	findAPRecursive(v)
	return articulationPoints
}

//problematic
func findAPRecursive(v *vertex) (lowDiscovery int) {
	discovery[v] = time
	lowerDiscovery[v] = time
	time++

	for _, av := range v.adjacent {
		if discovery[av] == 0 {
			// 0代表这个点还没有被访问
			predecessor[av] = v
			lowOfAv := findAPRecursive(av)
			if lowOfAv < lowerDiscovery[v] {
				lowerDiscovery[v] = lowOfAv
			}

		} else {
			// 这个点之前被访问过了
			if predecessor[v] != av && discovery[av] < lowerDiscovery[v] {
				// 第一个如果是true，则代表这是回边
				// 并且如果这个邻接点，包括前面的点，比我的能去的最先的小
				// 我的lower就可以设置更小的
				lowerDiscovery[v] = discovery[av]
			}
		}
	}

	if discovery[v] <= lowerDiscovery[v] || (predecessor[v] == nil && len(v.adjacent) > 1) {
		articulationPoints = append(articulationPoints, v)
	}

	return lowerDiscovery[v]
}

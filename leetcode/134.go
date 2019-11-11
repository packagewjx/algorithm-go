//+build 134me

package leetcode

import "container/heap"

type rest struct {
	restGas int
	pos     int
}

type restHeap []*rest

func (r *restHeap) Len() int {
	return len(*r)
}

func (r *restHeap) Less(i, j int) bool {
	return (*r)[i].restGas > (*r)[j].restGas
}

func (r *restHeap) Swap(i, j int) {
	temp := (*r)[i]
	(*r)[i] = (*r)[j]
	(*r)[j] = temp
}

func (r *restHeap) Push(x interface{}) {
	*r = append(*r, x.(*rest))
}

func (r *restHeap) Pop() interface{} {
	ret := (*r)[len(*r)-1]
	*r = (*r)[:len(*r)-1]
	return ret
}

func canCompleteCircuit(gas []int, cost []int) int {
	rh := make(restHeap, 0, len(gas))
	heap.Init(&rh)
	for i := 0; i < len(gas); i++ {
		heap.Push(&rh, &rest{
			restGas: gas[i] - cost[i],
			pos:     i,
		})
	}

	// 从最大的开始
	for i := 0; i < rh.Len(); i++ {
		r := heap.Pop(&rh).(*rest)
		if r.restGas < 0 {
			return -1
		}
		// 验证
		canGo := true
		g := 0
		for j := r.pos; j < len(cost); j++ {
			g += gas[j]
			if g < cost[j] {
				canGo = false
				break
			}
			g -= cost[j]
		}
		if canGo {
			for j := 0; j < r.pos; j++ {
				g += gas[j]
				if g < cost[j] {
					canGo = false
					break
				}
				g -= cost[j]
			}
			if canGo {
				return r.pos
			}
		}
	}

	return -1
}

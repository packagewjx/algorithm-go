package study

import (
	"sort"
	"sync"
)

type EditDistance struct {
	word     string
	distance int
}

func getDistanceDP(w1, w2 string, i, j int, memo [][]int) int {
	if i == len(w1) {
		return len(w2) - j
	} else if j == len(w2) {
		return len(w1) - i
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	// 实际计算编辑距离

	// 如果相等，则直接查看i+1和j+1即可
	if w1[i] == w2[j] {
		ret := getDistanceDP(w1, w2, i+1, j+1, memo)
		memo[i][j] = ret
		return ret
	}

	// 不相等的话，就要查看5种情况

	// 删除w1[i]或者在w2[j]前面加入字符w1[i]
	smallest := getDistanceDP(w1, w2, i+1, j, memo)
	// 删除w2[j]或者在w1[i]前面加入字符w2[j]
	temp := getDistanceDP(w1, w2, i, j+1, memo)
	if temp < smallest {
		smallest = temp
	}
	// 将w1[i]替换为w2[j]或者反过来
	temp = getDistanceDP(w1, w2, i+1, j+1, memo)
	if temp < smallest {
		smallest = temp
	}
	// 经过上述的操作之后，编辑距离加1
	smallest += 1
	memo[i][j] = smallest
	return smallest
}

// 计算w1与w2编辑距离的函数
func getEditDistance(w1 string, w2 string) int {
	memo := make([][]int, len(w1))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(w2))
		for j := 0; j < len(w2); j++ {
			memo[i][j] = -1
		}
	}
	return getDistanceDP(w1, w2, 0, 0, memo)
}

// 计算word与dict中的所有word的编辑距离。
// 编辑距离指的是，在字符串的一个位置上添加、删除、或修改字符，能够与另一个字符串的一个位置相同的编辑，此次修改为1个编辑距离。所有的编辑距离加起来，
// 取最小值，就是两个单词的编辑距离
func GetEditDistances(word string, dict []string) []*EditDistance {
	distances := make([]*EditDistance, len(dict))
	wg := &sync.WaitGroup{}
	for i := 0; i < len(dict); i++ {
		wg.Add(1)
		go func(index int) {
			distance := getEditDistance(word, dict[index])
			distances[index] = &EditDistance{
				word:     dict[index],
				distance: distance,
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	return distances
}

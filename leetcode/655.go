package leetcode

import "strconv"

func printTree(root *TreeNode) [][]string {
	midResult := make([][]string, 0, 10)
	lastLevel := -1
	type context struct {
		level int
		node  *TreeNode
	}

	queue := make([]*context, 1, 16)
	queue[0] = &context{
		level: 0,
		node:  root,
	}
	nextAllNil := false
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.level != lastLevel {
			// 如果这一层都是nil，则不需要继续了
			if nextAllNil {
				break
			}
			midResult = append(midResult, make([]string, 0, 16))
			lastLevel = cur.level
			nextAllNil = true
		}
		if cur.node == nil {
			midResult[cur.level] = append(midResult[cur.level], "")
			queue = append(queue, &context{
				level: cur.level + 1,
				node:  nil,
			}, &context{
				level: cur.level + 1,
				node:  nil,
			})
		} else {
			nextAllNil = false
			midResult[cur.level] = append(midResult[cur.level], strconv.Itoa(cur.node.Val))
			queue = append(queue, &context{
				level: cur.level + 1,
				node:  cur.node.Left,
			}, &context{
				level: cur.level + 1,
				node:  cur.node.Right,
			})
		}
	}
	midResult = midResult[:len(midResult)-1]

	mi := 1
	totalLength := 1<<uint(len(midResult)) - 1
	result := make([][]string, len(midResult))
	for i := len(midResult) - 1; i >= 0; i-- {
		levelResult := make([]string, 0, totalLength)
		paddingLength := mi - 1
		marginLength := mi<<1 - 1
		for j := 0; j < paddingLength; j++ {
			levelResult = append(levelResult, "")
		}
		for j := 0; j < len(midResult[i])-1; j++ {
			levelResult = append(levelResult, midResult[i][j])
			for k := 0; k < marginLength; k++ {
				levelResult = append(levelResult, "")
			}
		}
		levelResult = append(levelResult, midResult[i][len(midResult[i])-1])
		for j := 0; j < paddingLength; j++ {
			levelResult = append(levelResult, "")
		}
		result[i] = levelResult
		mi <<= 1
	}

	return result
}

//["", "", "", "", "", "", "", "1", "", "",  "", "", "", "", ""]
//["", "", "", "2", "", "", "", "", "",  "", "", "5", "", "", ""]
//["", "3", "", "", "", "", "", "",  "", "", "", "", "", "", ""]
//["4", "",  "", "", "", "", "", "", "", "", "", "",  "", "", ""]

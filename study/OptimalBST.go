package study

import "math"

// 计算最优二叉搜索树，输入为概率。输出为根数组。输出的最后一行和第一列是没有意义的。
// 如我需要求由0到2的元素的最优BST，需要找root[0][3]
func OptimalBST(possibility []float64) (root [][]int) {
	if len(possibility) == 0 {
		return nil
	} else if len(possibility) == 1 {
		return [][]int{{0, 1}, {0, 0}}
	}

	root = make([][]int, len(possibility)+1)
	sum := make([][]float64, len(possibility)+1)
	for i := 0; i < len(root); i++ {
		root[i] = make([]int, len(possibility)+1)
		sum[i] = make([]float64, len(possibility)+1)
	}

	// 在对角线填入初始值
	for i := 0; i < len(possibility); i++ {
		sum[i][i+1] = possibility[i]
		root[i][i+1] = i
	}

	// 对角线赋值与计算
	// 这里的ij从0开始计算
	for jStart := 1; jStart < len(possibility); jStart++ {
		i := 0
		for j := jStart; j < len(possibility); j++ {
			possibilitySum := float64(0)
			for id := i; id <= j; id++ {
				possibilitySum += possibility[id]
			}

			minVal := math.MaxFloat64
			minRoot := -1
			// 选取k
			for k := i; k <= j; k++ {
				// 第二个下标需要加1，因为第一列是无意义的
				val := sum[i][k-1+1] + sum[k+1][j+1]
				if val < minVal {
					minVal = val
					minRoot = k
				}
			}

			sum[i][j+1] = minVal + possibilitySum
			root[i][j+1] = minRoot

			// i加上一行
			i++
		}
	}

	return root
}

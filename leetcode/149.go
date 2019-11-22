package leetcode

import (
	"fmt"
	"strconv"
)

func maxPoints(points [][]int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}

	if len(points) <= 1 {
		return len(points)
	}
	max := 1
	for i := 0; i < len(points); i++ {
		maxI := 1
		dup := 0
		lines := make(map[string]int)
		for j := i + 1; j < len(points); j++ {
			xi, yi, xj, yj := points[i][0], points[i][1], points[j][0], points[j][1]
			// 重合点加入到dup中
			var key string
			if xi == xj && yi == yj {
				dup++
				continue
			} else if xi == xj {
				key = strconv.Itoa(xi)
			} else {
				// 更高精度的gcd
				slopeGCD := gcd(yi-yj, xi-xj)
				slopeFenmu := (yi - yj) / slopeGCD
				slopeFenzi := (xi - xj) / slopeGCD
				// 不需要计算截距，因为经过i点斜率确定的话，截距也是确定的
				key = fmt.Sprintf("%d/%d", slopeFenmu, slopeFenzi)
			}

			if lines[key] == 0 {
				lines[key] = 2
			} else {
				lines[key]++
			}
			if lines[key] > maxI {
				maxI = lines[key]
			}
		}
		if max < dup+maxI {
			max = dup + maxI
		}
	}

	return max
}

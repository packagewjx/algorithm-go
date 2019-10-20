package leetcode

func checkStraightLine(coordinates [][]int) bool {
	// 取两点计算斜截式，首先查看x是否相等
	if coordinates[0][0] == coordinates[1][0] {
		// 线为x = coordinates[0][0]
		for i := 2; i < len(coordinates); i++ {
			if coordinates[i][0] != coordinates[0][0] {
				return false
			}
		}
		return true
	} else {
		slope := (coordinates[1][1] - coordinates[0][1]) / (coordinates[1][0] - coordinates[0][0])
		intercept := coordinates[0][1] - slope*coordinates[0][0]

		// 从第三个点开始寻找
		for i := 2; i < len(coordinates); i++ {
			if coordinates[i][1] != slope*coordinates[i][0]+intercept {
				return false
			}
		}
		return true
	}
}

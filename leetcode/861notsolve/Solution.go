package main

var rowTotal, colTotal, total int

func matrixScore(A [][]int) int {
	rowTotal = len(A)
	colTotal = len(A[0])
	total = rowTotal + colTotal
	return recursiveTryFlip(A, 0)
}

func recursiveTryFlip(A [][]int, curPos int) int {
	// curPos指的是当前转还是不转的位置。若为0到行数-1，则是决定行，若是行数到总数减1，则决定列
	if curPos == total {
		// 已经到末尾，返回
		return calSum(A)
	}
	// 首先不翻转当前位置，进行下一个
	noFlip := recursiveTryFlip(A, curPos+1)

	// 然后我们翻转当前位置，再进入下一个
	var flip int
	Acopy := copyA(A)
	if curPos < rowTotal {
		flipRow(Acopy, curPos)
		flip = recursiveTryFlip(Acopy, curPos+1)
	} else {
		flipColumn(Acopy, curPos-rowTotal)
		flip = recursiveTryFlip(Acopy, curPos+1)
	}

	if flip > noFlip {
		return flip
	} else {
		return noFlip
	}

}

func flipRow(A [][]int, rowNum int) {
	row := A[rowNum]
	for i, val := range row {
		row[i] = 1 - val
	}
}

func flipColumn(A [][]int, colNum int) {
	for _, row := range A {
		row[colNum] = 1 - row[colNum]
	}
}

func calSum(A [][]int) int {
	sum := 0
	for _, row := range A {
		for i, val := range row {
			sum += val << uint(colTotal-i-1)
		}
	}
	return sum
}

func copyA(A [][]int) [][]int {
	Acopy := make([][]int, len(A))
	for i, row := range A {
		Acopy[i] = make([]int, len(A[i]))
		copy(Acopy[i], row)
	}
	return Acopy
}

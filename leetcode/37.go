package leetcode

type sudokuPos struct {
	x    int
	y    int
	nums []int
}

type sudokuQueue struct {
	queue    [][]*sudokuPos
	entryMap [][]*sudokuPos
	smallest int
}

func newSudokuQueue() *sudokuQueue {
	queue := &sudokuQueue{
		queue:    make([][]*sudokuPos, 10),
		entryMap: make([][]*sudokuPos, 9),
		smallest: 9}

	for i := 0; i < 9; i++ {
		queue.queue[i] = make([]*sudokuPos, 0, 16)
		queue.entryMap[i] = make([]*sudokuPos, 9)
	}
	queue.queue[9] = make([]*sudokuPos, 0, 16)
	return queue
}

func (q *sudokuQueue) push(pos *sudokuPos) {
	l := len(pos.nums)
	q.queue[l] = append(q.queue[l], pos)
	if l < q.smallest {
		q.smallest = l
	}
	q.entryMap[pos.x][pos.y] = pos
}

func (q *sudokuQueue) pop() *sudokuPos {
	if q.smallest == 10 {
		return nil
	}
	ret := q.queue[q.smallest][0]
	q.queue[q.smallest] = q.queue[q.smallest][1:]
	if len(q.queue[q.smallest]) == 0 {
		i := q.smallest + 1
		for ; i < 10; i++ {
			if len(q.queue[i]) > 0 {
				q.smallest = i
				break
			}
		}
		if i == 10 {
			// 说明没有了
			q.smallest = 10
		}
	}
	q.entryMap[ret.x][ret.y] = nil
	return ret
}

func (q *sudokuQueue) getEntry(x, y int) *sudokuPos {
	ret := q.entryMap[x][y]
	if ret == nil {
		return nil
	}
	posQueue := &q.queue[len(ret.nums)]
	for i := 0; i < len(*posQueue); i++ {
		if (*posQueue)[i] == ret {
			*posQueue = append((*posQueue)[:i], (*posQueue)[i+1:]...)
			break
		}
	}
	q.entryMap[x][y] = nil
	if len(*posQueue) == 0 {
		// 空了，要设置新的
		i := q.smallest + 1
		for ; i < 10; i++ {
			if len(q.queue[i]) > 0 {
				q.smallest = i
				break
			}
		}
		if i == 10 {
			// 说明没有了
			q.smallest = 10
		}
	}
	return ret
}

func (q *sudokuQueue) saveEntry(pos *sudokuPos) {
	q.queue[len(pos.nums)] = append(q.queue[len(pos.nums)], pos)
	q.entryMap[pos.x][pos.y] = pos
	if len(pos.nums) < q.smallest {
		q.smallest = len(pos.nums)
	}
}

func deleteEntryNum(num int, p *sudokuPos) bool {
	for i := 0; i < len(p.nums); i++ {
		if p.nums[i] == num {
			p.nums = append(p.nums[:i], p.nums[i+1:]...)
			return true
		}
	}
	return false
}

func solveSudokuRecursive(board [][]byte, q *sudokuQueue) bool {
	pos := q.pop()
	if pos == nil {
		// 没有了，返回true
		return true
	}

	// 将需要修改的所有的加入队列中
	toChangePos := make([][]int, 0, 16)
	for j := 0; j < 9; j++ {
		if board[j][pos.y] != '.' || j == pos.x {
			continue
		}
		toChangePos = append(toChangePos, []int{j, pos.y})
	}
	for j := 0; j < 9; j++ {
		if board[pos.x][j] != '.' || j == pos.y {
			continue
		}
		toChangePos = append(toChangePos, []int{pos.x, j})
	}
	for j := pos.x / 3 * 3; j < pos.x/3*3+3; j++ {
		for k := pos.y / 3 * 3; k < pos.y/3*3+3; k++ {
			if j == pos.x || k == pos.y || board[j][k] != '.' {
				continue
			}
			toChangePos = append(toChangePos, []int{j, k})
		}
	}

	for i := 0; i < len(pos.nums); i++ {
		num := pos.nums[i]
		deleted := make([]bool, len(toChangePos))
		for j := 0; j < len(toChangePos); j++ {
			entry := q.getEntry(toChangePos[j][0], toChangePos[j][1])
			deleted[j] = deleteEntryNum(num, entry)
			q.saveEntry(entry)
		}

		//填入
		board[pos.x][pos.y] = '0' + byte(num) + 1
		if solveSudokuRecursive(board, q) {
			return true
		} else {
			// 复原
			for j := 0; j < len(toChangePos); j++ {
				if deleted[j] {
					entry := q.getEntry(toChangePos[j][0], toChangePos[j][1])
					entry.nums = append(entry.nums, num)
					q.saveEntry(entry)
				}
			}
		}
	}
	// 全部数字都不行的话，就复原
	board[pos.x][pos.y] = '.'
	q.push(pos)
	return false
}

func solveSudoku(board [][]byte) {
	q := newSudokuQueue()
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	grid := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 9)
		col[i] = make([]bool, 9)
		grid[i] = make([]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '0' - 1
			row[i][num] = true
			col[j][num] = true
			grid[i/3*3+j/3][num] = true
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			pos := &sudokuPos{
				x:    i,
				y:    j,
				nums: []int{},
			}
			for num := 0; num < 9; num++ {
				if (row[i][num] || col[j][num] || grid[i/3*3+j/3][num]) == false {
					pos.nums = append(pos.nums, num)
				}
			}
			q.push(pos)
		}
	}

	solveSudokuRecursive(board, q)
}

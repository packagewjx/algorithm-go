package leetcode

import "strconv"

// LeetCode提交不通过，但是结果是对的。。。

func stonePileTake(piles []int, M int, X int, start int) (thisTake, newM int) {
	if X < 0 || X > 2*M {
		print("error")
		return 0, 0
	}
	thisTake = 0
	for i := 0; i < X && i+start < len(piles); i++ {
		thisTake += piles[start+i]
	}
	if X > M {
		newM = X
	} else {
		newM = M
	}
	return
}

var stoneGameMemo = make(map[string]*memos)

type memos struct {
	take int
	x    int
	newM int
}

func stoneGameIIDP(piles []int, start int, M int, meFirst bool) (take, x, newM int) {
	if start < 0 || start >= len(piles) {
		return 0, 0, M
	}
	// 如果M过大，没有意义的
	upperBound := 2 * M
	if len(piles)-start < 2*M {
		upperBound = len(piles) - start
	}

	key := strconv.Itoa(start) + ":" + strconv.Itoa(upperBound)
	val, ok := stoneGameMemo[key]
	if ok {
		if meFirst {
			return val.take, val.x, val.newM
		} else {
			return stoneGameIIDP(piles, start+val.x, val.newM, false)
		}
	}

	biggest := memos{
		take: 0,
		x:    0,
		newM: 0,
	}
	if meFirst {
		for i := 1; i <= upperBound; i++ {
			thisTake, m := stonePileTake(piles, M, i, start)
			iTake, _, _ := stoneGameIIDP(piles, start+i, m, false)
			totalTake := thisTake + iTake
			if totalTake > biggest.take {
				biggest.take = totalTake
				biggest.x = i
				biggest.newM = m
			}
		}
	} else {
		_, x2, m2 := stoneGameIIDP(piles, start, M, true)
		i, x3, m3 := stoneGameIIDP(piles, start+x2, m2, true)
		biggest.take = i
		biggest.x = x3
		biggest.newM = m3
	}
	if meFirst {
		stoneGameMemo[key] = &biggest
	}

	return biggest.take, biggest.x, biggest.newM
}

func stoneGameII(piles []int) int {
	take, _, _ := stoneGameIIDP(piles, 0, 1, true)
	return take
}

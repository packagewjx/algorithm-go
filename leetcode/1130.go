package leetcode

import (
	"math"
)

type mctResult struct {
	maxInt int
	sum    int
}

var mctMemo [][]*mctResult

func mctFromLeafValuesDP(arr []int, begin, end int) (maxInt, sum int) {
	if begin >= end {
		return 0, 0
	}
	if begin == end-1 {
		return arr[begin], 0
	}
	result := mctMemo[begin][end]
	if result != nil {
		return result.maxInt, result.sum
	}

	r := &mctResult{
		maxInt: math.MinInt64,
		sum:    math.MaxInt64,
	}
	for i := begin + 1; i < end; i++ {
		leftMaxInt, leftSum := mctFromLeafValuesDP(arr, begin, i)
		rightMaxInt, rightSum := mctFromLeafValuesDP(arr, i, end)
		sum := leftSum + rightSum + leftMaxInt*rightMaxInt
		if sum < r.sum {
			r.sum = sum
			if leftMaxInt < rightMaxInt {
				r.maxInt = rightMaxInt
			} else {
				r.maxInt = leftMaxInt
			}
		}
	}
	mctMemo[begin][end] = r
	return r.maxInt, r.sum
}

func mctFromLeafValues(arr []int) int {
	mctMemo = make([][]*mctResult, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		mctMemo[i] = make([]*mctResult, len(arr)+1)
	}
	_, sum := mctFromLeafValuesDP(arr, 0, len(arr))
	return sum
}

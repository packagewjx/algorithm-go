package leetcode

import "math"

func minHeightShelvesDP(books [][]int, shelfWidth int, cur int, memo []*int) int {
	if memo[cur] != nil {
		return *memo[cur]
	}

	minHeight := new(int)
	*minHeight = math.MaxInt64
	width := books[cur][0]
	height := books[cur][1]
	bookCount := 1
	for cur+bookCount < len(books) && width <= shelfWidth {
		totalHeight := minHeightShelvesDP(books, shelfWidth, cur+bookCount, memo)
		if height+totalHeight < *minHeight {
			*minHeight = height + totalHeight
		}
		width += books[cur+bookCount][0]
		if books[cur+bookCount][1] > height {
			height = books[cur+bookCount][1]
		}
		bookCount++
	}
	// 处理最后一本的情况
	if cur+bookCount == len(books) && width <= shelfWidth {
		if height < *minHeight {
			*minHeight = height
		}
	}

	memo[cur] = minHeight
	return *minHeight
}

func minHeightShelves(books [][]int, shelf_width int) int {
	memo := make([]*int, len(books))
	memo[len(books)-1] = &books[len(books)-1][1]
	return minHeightShelvesDP(books, shelf_width, 0, memo)
}

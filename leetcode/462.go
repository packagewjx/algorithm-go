package leetcode

func partition(nums []int, start int, end int) int {
	partedEnd := start
	pivot := nums[end-1]
	for i := start; i < end; i++ {
		if nums[i] < pivot {
			temp := nums[partedEnd]
			nums[partedEnd] = nums[i]
			nums[i] = temp
			partedEnd++
		}
	}
	nums[end-1] = nums[partedEnd]
	nums[partedEnd] = pivot
	return partedEnd
}

func findKthNum(nums []int, K int) int {
	p := 0
	q := len(nums)
	m := partition(nums, p, q)
	for m != K-1 {
		if m < K-1 {
			p = m + 1
		} else {
			q = m
		}
		m = partition(nums, p, q)
	}
	return nums[m]
}

func minMoves2(nums []int) int {
	median := 0
	if len(nums)&1 == 1 {
		median = findKthNum(nums, len(nums)/2+1)
	} else {
		num1 := findKthNum(nums, len(nums)/2)
		num2 := findKthNum(nums, len(nums)/2+1)
		median = (num1 + num2) / 2
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > median {
			sum += nums[i] - median
		} else {
			sum += median - nums[i]
		}
	}
	return sum
}

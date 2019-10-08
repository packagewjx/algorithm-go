package leetcode

func longestSubsequence(arr []int, difference int) int {
	res := 1
	length := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		last := length[arr[i]-difference]
		this := length[arr[i]]
		if last+1 > this {
			length[arr[i]] = last + 1
			if last+1 > res {
				res = last + 1
			}
		}
	}

	return res
}

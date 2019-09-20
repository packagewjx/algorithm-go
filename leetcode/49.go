package leetcode

func groupAnagrams(strs []string) [][]string {
	primes := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107}

	// 还是有可能发生冲突的。冲突发生时候需要比对，并且会遇到性能退化的情况
	set := make(map[int][]string)
	for i := 0; i < len(strs); i++ {
		h := 1
		for j := 0; j < len(strs[i]); j++ {
			h *= primes[strs[i][j]&31]
		}
		set[h] = append(set[h], strs[i])
	}

	result := make([][]string, 0, len(set))
	for _, v := range set {
		result = append(result, v)
	}
	return result
}

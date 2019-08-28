package leetcode

func beautifulArray(N int) []int {
	res := []int{1}
	for len(res) < N {
		tmp := make([]int, 0, 2*len(res))
		for i := 0; i < len(res); i++ {
			if res[i]*2-1 <= N {
				tmp = append(tmp, res[i]*2-1)
			}
		}
		for i := 0; i < len(res); i++ {
			if res[i]*2 <= N {
				tmp = append(tmp, res[i]*2)
			}
		}
		res = tmp
	}
	return res
}

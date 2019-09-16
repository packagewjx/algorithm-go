package leetcode

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	letters := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	lastResults := []string{""}
	for i := 0; i < len(digits); i++ {
		res := make([]string, 0, len(lastResults)*4)
		digit := digits[i] & 15
		for j := 0; j < len(lastResults); j++ {
			for k := 0; k < len(letters[digit]); k++ {
				res = append(res, lastResults[j]+letters[digit][k:k+1])
			}
		}
		lastResults = res
	}
	return lastResults
}

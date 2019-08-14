package leetcode

import "strings"

func findDuplicate(paths []string) [][]string {
	groups := make(map[string][]string)

	for i := 0; i < len(paths); i++ {
		path := paths[i]
		split := strings.Split(path, " ")
		directory := split[0]
		for j := 1; j < len(split); j++ {
			leftPerIndex := strings.Index(split[j], "(")
			content := split[j][leftPerIndex+1 : len(split[j])-1]
			fileName := directory + "/" + split[j][:leftPerIndex]
			group, ok := groups[content]
			if !ok {
				group = make([]string, 0, 10)
			}
			group = append(group, fileName)
			groups[content] = group
		}
	}

	result := make([][]string, 0, len(groups))
	for _, files := range groups {
		if len(files) > 1 {
			result = append(result, files)
		}
	}
	return result
}

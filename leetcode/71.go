package leetcode

func simplifyPath(path string) string {
	pathName := make([]string, 0, 10)
	handlePath := func(dirName string) {
		switch dirName {
		case ".":
			return
		case "..":
			if len(pathName) > 0 {
				pathName = pathName[:len(pathName)-1]
			}
		default:
			pathName = append(pathName, dirName)
		}
	}

	start := 0
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			if i > start {
				handlePath(path[start:i])
			}
			start = i + 1
		}
	}
	// 最后的处理
	if start < len(path) {
		handlePath(path[start:])
	}

	if len(pathName) == 0 {
		return "/"
	} else {
		res := ""
		for i := 0; i < len(pathName); i++ {
			res += "/" + pathName[i]
		}
		return res
	}
}

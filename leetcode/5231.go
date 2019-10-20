package leetcode

import "strings"

type my5231Node struct {
	name string
	next map[string]*my5231Node
	end  bool
}

func my5231DFS(node *my5231Node, cur string, result *[]string) {
	if node.end == true {
		*result = append(*result, cur)
	}

	cur += "/"
	for name, n := range node.next {
		my5231DFS(n, cur+name, result)
	}
}

func removeSubfolders(folder []string) []string {
	head := &my5231Node{
		name: "",
		next: make(map[string]*my5231Node),
		end:  false,
	}

	for _, fd := range folder {
		name := strings.Split(fd, "/")
		cur := head
		for i := 1; i < len(name); i++ {
			if cur.next[name[i]] == nil {
				newNode := &my5231Node{
					name: name[i],
					next: make(map[string]*my5231Node),
				}
				cur.next[name[i]] = newNode
			}
			cur = cur.next[name[i]]
			if cur.end {
				break
			}
		}

		if cur.end {
			continue
		}
		cur.end = true
		// 清除子文件夹
		cur.next = nil
	}

	// 遍历树
	result := make([]string, 0, 10)
	my5231DFS(head, "", &result)
	return result
}

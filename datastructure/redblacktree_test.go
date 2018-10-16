package datastructure

import (
	"fmt"
	"testing"
)

func TestRedBlackTree_Insert(t *testing.T) {
	nums := []int{12, 1, 9, 2, 0, 11, 7, 19, 4, 15, 18, 5, 14, 13, 10, 16, 6, 3, 8, 17}

	tree := &RedBlackTree{}
	for _, value := range nums {
		tree.Insert(&dummy{key: value})
	}
	printTree(tree)
}

func TestRedBlackTree_Delete(t *testing.T) {
	nums := []int{12, 1, 9, 2, 0, 11, 7, 19, 4, 15, 18, 5, 14, 13, 10, 16, 6, 3, 8, 17}

	tree := &RedBlackTree{}
	for _, value := range nums {
		tree.Insert(&dummy{key: value})
	}

	for _, value := range nums {
		tree.Delete(value)
	}

}

func printTree(tree *RedBlackTree) {
	if tree.root == nil {
		return
	}

	type temp struct {
		node  *rbNode
		level int
	}

	queue := make([]*temp, 0, 10)
	queue = append(queue, &temp{tree.root, 1})
	currentLevel := 0
	for len(queue) > 0 {
		node := queue[0]
		if currentLevel != node.level {
			currentLevel = node.level
			fmt.Println()
			fmt.Print(currentLevel, ":")
		}

		if node.node != nil {
			var color string
			if node.node.color == RED {
				color = "R"
			} else {
				color = "B"
			}
			fmt.Printf(" %d%s ", node.node.key(), color)

			if !(node.node.leftChild == nil && node.node.rightChild == nil) {
				queue = append(queue, &temp{node.node.leftChild, currentLevel + 1})
				queue = append(queue, &temp{node.node.rightChild, currentLevel + 1})
			}
		} else {
			fmt.Print("<nil>B")
		}

		queue = queue[1:]
	}
	fmt.Println()

}

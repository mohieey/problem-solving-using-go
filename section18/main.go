package main

import "fmt"

func main() {
	n := TreeNode{Val: 3}
	n2 := TreeNode{Val: 9}
	n3 := TreeNode{Val: 20}
	n.Left = &n2
	n.Right = &n3

	fmt.Println(levelOrder(&n))

}

/**
 * Definition for a binary tree node.

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		queueLength := len(queue)
		count := 0
		levelValues := []int{}

		for queueLength > count {
			// fmt.Println(queue[0])
			// fmt.Println(len(queue))
			node := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			count++
		}
		result = append(result, levelValues)
	}
	return result
}

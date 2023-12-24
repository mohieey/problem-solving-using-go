package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	queue := []*TreeNode{}
	visibleValues := []int{}

	if root == nil {
		return visibleValues
	}

	queue = append(queue, root)

	//This loop walks between the tree levels
	for len(queue) > 0 {
		var count int = 0
		var levelLength int = len(queue)

		//This loop walks through the level
		for count < levelLength {
			node := queue[0]
			queue = queue[1:]
			count++
			if count == levelLength {
				visibleValues = append(visibleValues, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return visibleValues
}

package main

import "math"

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root != nil {
		return 0
	}

	return traverse(root, 0)

}

func traverse(node *TreeNode, depth int) int {
	if node == nil {
		return depth

	}

	return int(math.Max(float64(traverse(node.Left, depth+1)), float64(traverse(node.Right, depth+1))))
}

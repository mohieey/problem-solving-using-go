package main

import (
	"fmt"
	"math"
)

func main() {

	n1 := TreeNode{Val: 1}
	n2 := TreeNode{Val: 2}
	n3 := TreeNode{Val: 3}
	n2.Left = &n1
	n2.Right = &n3
	fmt.Println(isValidBST(&n2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validateNode(root, int(math.Inf(-1)), int(math.Inf(1)))
}

func validateNode(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if !(min < node.Val && node.Val < max) {
		return false
	}

	return validateNode(node.Left, min, node.Val) && validateNode(node.Right, node.Val, max)
}

package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	nexts := []*Node{}

	var prevNode, currentNode *Node = nil, root

	for currentNode != nil {
		if currentNode.Child != nil {
			if currentNode.Next != nil {
				nexts = append(nexts, currentNode.Next)
			}
			child := currentNode.Child
			child.Prev = currentNode
			currentNode.Next = child
			currentNode.Child = nil
		}
		prevNode = currentNode
		currentNode = currentNode.Next
		if currentNode == nil && len(nexts) != 0 {
			currentNode = nexts[len(nexts)-1]
			prevNode.Next = currentNode
			currentNode.Prev = prevNode
			nexts = nexts[:len(nexts)-1]
		}
	}

	return root
}

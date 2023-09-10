package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right || head == nil {
		return head
	}
	var prevLeftNode, leftNode, nextRightNode, rightNode *ListNode = nil, nil, nil, nil

	var prevNode, currentNode *ListNode = nil, head

	var i = 1

	for i <= left {
		if i == left {
			prevLeftNode = prevNode // nil
			leftNode = currentNode  // 3
		}
		prevNode = currentNode         // 3
		currentNode = currentNode.Next // 5
		i++                            // 2
	}

	for i <= right {
		if i == right {
			nextRightNode = currentNode.Next // nil
			rightNode = currentNode          // 5
		}
		nextNode := currentNode.Next // nil
		currentNode.Next = prevNode  // 3
		prevNode = currentNode       // 5
		currentNode = nextNode       // nil
		i++
	}

	if left > 1 {
		prevLeftNode.Next = rightNode
	}
	leftNode.Next = nextRightNode

	if left == 1 {
		return rightNode
	}

	if nextRightNode == nil && left == 1 {
		return rightNode
	}

	return head

}

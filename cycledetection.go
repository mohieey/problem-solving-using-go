package main

func detectCycle(head *ListNode) *ListNode {
	var tortoise, hare *ListNode = head, head

	for hare != nil {
		hare = hare.Next

		if hare == nil || hare.Next == nil {
			return nil
		}

		hare = hare.Next
		tortoise = tortoise.Next

		if tortoise == hare {
			var p1, p2 *ListNode = head, hare
			for {
				if p1 == p2 {
					return p1
				}
				p1 = p1.Next
				p2 = p2.Next
			}
		}

	}

	return nil
}

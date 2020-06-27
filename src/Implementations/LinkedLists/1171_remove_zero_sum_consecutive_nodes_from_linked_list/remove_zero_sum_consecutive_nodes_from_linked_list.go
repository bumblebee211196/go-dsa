package remove_zero_sum_consecutive_nodes_from_linked_list

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Given the head of a linked list, we repeatedly delete consecutive
// sequences of nodes that sum to 0 until there are no such sequences.
// After doing so, return the head of the final linked list. You may return any such answer.

func RemoveZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	head = &ListNode{Val: -1, Next: head}
	curr := head
	for curr != nil {
		psum := 0
		next := curr.Next
		for next != nil {
			psum += next.Val
			if psum == 0 {
				curr.Next = next.Next
			}
			next = next.Next
		}
		curr = curr.Next
	}
	return head.Next
}

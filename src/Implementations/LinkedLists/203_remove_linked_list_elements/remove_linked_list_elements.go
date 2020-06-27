package remove_linked_list_elements

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Remove all elements from a linked list of integers that have value val.

// Solution: Straight forward approach
// We use tow pointer, one for the current node and other for the previous node.
// We run an outer loop traversing the entire list. While traversing each node, we run an
// inner loop which runs on a condition that the given node's Val equals the given val.
// When the inner loop finished, we point the previous' Next to the current's Next.
//
// Time Complexity: O(n)	Space Complexity: O(1)
func RemoveElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{Val: -1, Next: head}
	prev, curr := newHead, newHead.Next
	for curr != nil {
		for curr != nil && curr.Val == val {
			curr = curr.Next
		}
		prev.Next = curr
		prev = prev.Next
		if curr != nil {
			curr = curr.Next
		}
	}
	return newHead.Next
}

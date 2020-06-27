package swap_nodes_in_pairs

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Given a linked list, swap every two adjacent nodes and return its head.
// You may not modify the values in the list's nodes, only nodes itself may be changed.

// Solution: Straight forward approach
// We traverse the list by two nodes at a time, and we swap the two nodes by changing their
// Next pointer. We do this until we reach the end of the list.
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head = &ListNode{Val: -1, Next: head}
	curr := head
	for curr.Next != nil && curr.Next.Next != nil {
		first, second := curr.Next, curr.Next.Next
		tmp := second.Next
		second.Next, first.Next = first, tmp
		curr.Next = second
		curr = curr.Next.Next
	}
	return head.Next
}

package MiddleOfLinkedList

import . "github.com/bumblebee211996/go-ds-daa/src/Implementations/LinkedLists"

// Problem Statement: Find the middle node of a linked list. If there are two middle nodes,
// return the second middle node.

// Solution 1: Using the length of the linked list.
// We can calculate the length of the linked list and traverse the list till half of the length
// value. The node at the half length is the answer
func MidOfLinkedList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	l, curr := 0, head
	for curr != nil {
		l++
		curr = curr.Next
	}
	curr = head
	for i := 0; i < l/2; i++ {
		curr = curr.Next
	}
	return curr
}

// Solution 2: Using Flyod's algorithm.
// We will two pointers: slow and fast. The faster pointer traverses twice the speed of the slow
// pointer. When the fast pointer reaches the end, the node which the slow pointer points is the
// answer.
func MidOfLinkedList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

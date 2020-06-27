package reverse_linked_list_2

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Reverse a linked list from position m to n. Do it in one-pass.

// Solution: Straight forward approach
// We perform the normal reverse solution to the nodes in the linked list between the
// range m and n
func ReverseBetween(head *ListNode, m, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 1; i < m; i++ {
		prev = prev.Next
		head = head.Next
	}
	mth, nth := head, head.Next
	for i := m; i < n; i++ {
		nth.Next, head, nth = head, nth, nth.Next
	}
	mth.Next, prev.Next = nth, head
	return dummy.Next
}

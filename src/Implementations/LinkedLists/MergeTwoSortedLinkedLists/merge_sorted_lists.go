package MergeTwoSortedLinkedLists

import . "github.com/bumblebee211996/go-ds-daa/src/Implementations/LinkedLists"

// Problem Statement: Merge two sorted linked lists and return it as a new sorted list.

// Solution: WE will have two pointers each pointing to the head of the two lists respectively.
// We will create a new list and keep adding the nodes to this list. If the value of the first
// pointer is less than the second one, we add that node to the new list and move forward that
// pointer. Similar is the case for the other pointer. We do this until any one of them reaches
// the end. When one fo the list reaches the end, we just the append the remaining part(if left any)
// of the other list to the new list and return the head of the new list.
func MergeSortedLinkedLists(head1, head2 *ListNode) *ListNode {
	newHead := &ListNode{Val: -1}
	curr := newHead
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			curr.Next = &ListNode{Val: head1.Val}
			head1, curr = head1.Next, curr.Next
		} else {
			curr.Next = &ListNode{Val: head2.Val}
			head2, curr = head2.Next, curr.Next
		}
	}
	for head1 != nil {
		curr.Next = &ListNode{Val: head1.Val}
		head1, curr = head1.Next, curr.Next
	}
	for head2 != nil {
		curr.Next = &ListNode{Val: head2.Val}
		head2, curr = head2.Next, curr.Next
	}
	return newHead.Next
}

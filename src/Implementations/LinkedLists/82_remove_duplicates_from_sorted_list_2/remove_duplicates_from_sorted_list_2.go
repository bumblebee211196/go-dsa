package remove_duplicates_from_sorted_list_2

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Given a sorted linked list, delete all nodes that have duplicate numbers,
// leaving only distinct numbers from the original list.
//
// Return the linked list sorted as well.

// Solution: Straight forward approach
// We use two pointers prev and curr pointing to previous and current node respectively. We use
// two loops, th outer loop traverses until we reach the end of the list. The inner loop traverses
// with the condition that the next node's Val equals the current node's val. This way we reach
// the node which is next to the nodes with the same values. We simply make previous node's Next
// point to that node.
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr, newHead := head, &ListNode{Next: head}
	prev := newHead
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			val := curr.Val
			for curr != nil && val == curr.Val {
				curr = curr.Next
			}
			prev.Next = curr
		} else {
			prev = prev.Next
			curr = curr.Next
		}
	}
	return newHead.Next
}

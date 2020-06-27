package remove_duplicates_from_sorted_list

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Delete all duplicates such that each element appear only once.

// Solution: Straight forward approach
// We can determine if a node is a duplicate by comparing its value to the node after
// it in the list. If it is a duplicate, we change the next pointer of the current node
// so that it skips the next node and points directly to the one after the next node.
//
// Time Complexity: O(n)	Space Complexity: O(1)
func RemoveDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return dummy
}

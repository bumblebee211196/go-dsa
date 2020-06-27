package remove_nth_node_from_end_of_list

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Given a linked list, remove the n-th node from the end of list and return its head.

// Solution: Straight forward approach
// We use two pointers. With one pointer, we traverse the list till the given n value.
// Now with the other pointer we traverse the list until the first pointer reaches the end.
// When the fast pointer reaches the end, we reach our desired node.
//
// Time Complexity: O(n)	Space Complexity: O(1)
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		head = head.Next
	} else {
		for fast.Next != nil {
			slow = slow.Next
			fast = fast.Next
		}
		slow.Next = slow.Next.Next
	}
	return head
}

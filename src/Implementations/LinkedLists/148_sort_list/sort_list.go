package sort_list

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Sort a linked list in O(n log n) time using constant space complexity.

// Solution: Merge Sort
// We can find the mid of the linked list using Floyd's algorithm.
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// Finding mid node
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// Dividing the list into two parts
	l1, l2 := head, slow.Next
	slow.Next = nil
	// Split the lists again
	l1, l2 = SortList(l1), SortList(l2)
	// Merge the two sorted lists
	node := &ListNode{}
	head = node
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return head.Next
}

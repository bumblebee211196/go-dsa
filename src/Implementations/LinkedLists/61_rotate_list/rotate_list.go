package rotate_list

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Given a linked list, rotate the list to the right by k places,
// where k is non-negative.

// Solution: Straight forward approach
// We make the list circular.
// Original list: 1 --> 2 --> 3 --> 4 --> 5
//
// Circular list: 1 --> 2 --> 3 --> 4 --> 5
//               / \                      |
//                |_______________________|
//
// We then traverse this circular list till (n - k%n) times where n is the length of the list
// For example if k is 2, and if we traverse our current node will be at 3, we point current node's
// Next to nil and make the next node as head.
//
// Time Complexity: O(n)	Space Complexity: O(1)
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n, curr := 1, head
	for ; curr.Next != nil; curr = curr.Next {
		n++
	}
	if k%n == 0 {
		return head
	}
	curr.Next = head
	for i := 1; i < (n - k%n); i++ {
		head = head.Next
	}
	curr = head.Next
	head.Next = nil
	return curr
}

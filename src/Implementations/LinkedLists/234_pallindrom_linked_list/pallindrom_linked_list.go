package pallindrom_linked_list

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Check if the given linked list is a palindrome or no

// Solution: We find the mid point of the linked list. Break the list into
// two parts. Reverse one of the linked list and compare both of them.
// We can use Floyd's algorithm to find the mid point of the list.
//
// Time Complexity: O(n)	Space Complexity: O(1)
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	slow = reverse(slow)
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow, fast = slow.Next, fast.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

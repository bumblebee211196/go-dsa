package add_two_numbers

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: You are given two non-empty linked lists representing two
// non-negative integers. The digits are stored in reverse order and each of their
// nodes contain a single digit. Add the two numbers and return it as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Solution: Straight forward approach
//
// Time Complexity: O(n)	Space Complexity: O(n)
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	curr := &ListNode{Val: -1}
	head := curr
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		carry = sum / 10
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return head.Next
}

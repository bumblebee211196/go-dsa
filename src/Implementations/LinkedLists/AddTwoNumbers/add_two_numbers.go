package AddTwoNumbers

import (
	. "github.com/bumblebee211996/go-ds-daa/src/Implementations/LinkedLists"
	. "github.com/bumblebee211996/go-ds-daa/src/Implementations/LinkedLists/Reverse"
)

// Problem Statement: Add the two numbers and return it as a linked list. The most
// significant digit comes first and each of their nodes contain a single digit.

// Solution1: Using reverse
// We reverse both the lists and add the elements from both the lists one by one.
//
// Time Complexity: O(n)	Space Complexity: O(1)
func AddTwoNumbers1(head1, head2 *ListNode) *ListNode {
	head1, head2 = ReverseLinkedList2(head1), ReverseLinkedList2(head2)
	head := &ListNode{Val: -1}
	curr := head
	sum, carry := 0, 0
	for head1 != nil || head2 != nil {
		sum = carry
		if head1 != nil {
			sum += head1.Val
			head1 = head1.Next
		}
		if head2 != nil {
			sum += head2.Val
			head2 = head2.Next
		}
		curr.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return ReverseLinkedList2(head.Next)
}

// Solution 2: Using slices
// In the previous, it required to modify the list(reverse). With slices we
// create replicas of the lists with just values in it and do the addition
// and turn the result slice back into a list.
//
// Time Complexity: O(n)	Space Complexity: O(1)
func AddTwoNumbers2(head1, head2 *ListNode) *ListNode {
	l1, l2 := listToSlice(head1), listToSlice(head2)
	carry, sum := 0, 0
	var res *ListNode
	n1, n2 := len(l1), len(l2)
	for i := 0; i < n1 || i < n2; i++ {
		sum = carry
		if i < n1 {
			sum += l1[i]
		}
		if i < n2 {
			sum += l2[i]
		}
		res = &ListNode{Val: sum % 10, Next: res}
		carry = sum / 10
	}
	if carry > 0 {
		res = &ListNode{Val: carry, Next: res}
	}
	return res
}

func listToSlice(head *ListNode) []int {
	values := make([]int, 0)
	for curr := head; curr != nil; curr = curr.Next {
		values = append([]int{curr.Val}, values...)
	}
	return values
}

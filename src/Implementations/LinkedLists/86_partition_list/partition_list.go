package partition_list

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Given a linked list and a value x, partition it such that all
// nodes less than x come before nodes greater than or equal to x.
// You should preserve the original relative order of the nodes in each of the two partitions.

// Solution: Using two pointers
// We create a two new lists, one with values less than x and the other with values
// greater than or equal to x. Then we join the two lists. We are just reusing the same list,
// hence we don't use any extra space.
//
// Time Complexity: O(n)	Space Complexity: O(1)
func Partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstHead, secondHead := &ListNode{Val: -1}, &ListNode{Val: -1}
	first, second := firstHead, secondHead
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val < x {
			first.Next = curr
			first = first.Next
		} else {
			second.Next = curr
			second = second.Next
		}
	}
	second.Next = nil
	first.Next = secondHead.Next
	return firstHead.Next
}

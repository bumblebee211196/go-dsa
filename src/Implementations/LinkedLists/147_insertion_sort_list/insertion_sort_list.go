package insertion_sort_list

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Sort a linked list using insertion sort.

// Solution: Straight forward insertion sort
func InsertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := new(ListNode)
	newHead.Next = &ListNode{Val: head.Val}
	head = head.Next
	for ; head != nil; head = head.Next {
		prev, curr := newHead, newHead.Next
		for curr != nil {
			if head.Val < curr.Val {
				prev.Next = &ListNode{Val: head.Val, Next: curr}
				break
			}
			curr, prev = curr.Next, prev.Next
		}
		if curr == nil && head.Val >= prev.Val {
			prev.Next = &ListNode{Val: head.Val}
		}
	}
	return newHead.Next
}

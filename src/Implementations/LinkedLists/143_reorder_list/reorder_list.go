package reorder_list

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Given a singly linked list L: L0→L1→…→Ln-1→Ln,
// reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
//
// You may not modify the values in the list's nodes, only nodes itself may be changed.

// Solution: Straight forward approach
// We divide the list into two equal halves. Then we reverse the second half.
//
// Original List: 				L0 → L1 → L2 → ... → Ln-1 → Ln
//
// After splitting and reverse: L0 → L1 → ... Ln/2-1 → Ln/2		Ln → Ln-1 → ... → Ln/2-1 → Ln/2
//
// Now we create a new list by join ith node from both the lists.
//
// Time Complexity: O()		Space Complexity: O(1)
func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	} else {
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		left, right := head, slow.Next
		slow.Next = nil
		right = reverse(right)
		for left != nil && right != nil {
			lNext, rNext := left.Next, right.Next
			left.Next, right.Next = right, lNext
			left, right = lNext, rNext
		}
	}
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

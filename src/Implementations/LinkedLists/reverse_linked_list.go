package LinkedLists

// Problem statement: Reverse a linked list

// Solution 1:
// Create a new list by traversing each element and adding it to the
// head of the new list
// Time Complexity: O(n)	Space Complexity: O(n)
func ReverseLinkedList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead *ListNode
	for head != nil {
		newHead = &ListNode{head.Val, newHead}
		head = head.Next
	}
	return newHead
}

// Solution 2:
// We do in-place node reversals
// At first, we create a dummy head which is nil. Then traverse the list and
// make the node's Next point to the newHead and the node becomes the new head.
// Time Complexity: O(n)	Space Complexity: O(1)
func ReverseLinkedList2(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

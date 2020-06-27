package delete_node_in_linked_list

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Write a function to delete a node (except the tail) in a
// singly linked list, given only access to that node.

// Solution: Straight forward approach
// We update the current node with the next node's value, and point the current's
// "next" to the next node's "next". Since the node to be deleted will not be the
// tail node, we need not worry about nil node.
//
// Time Complexity: O(1)	Space Complexity: O(1)
func DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

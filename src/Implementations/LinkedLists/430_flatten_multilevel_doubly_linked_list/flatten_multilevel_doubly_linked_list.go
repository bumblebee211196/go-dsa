package flatten_multilevel_doubly_linked_list

// Problem Statement: You are given a doubly linked list which in addition to the
// next and previous pointers, it could have a child pointer, which may or may not
// point to a separate doubly linked list. These child lists may have one or more
// children of their own, and so on, to produce a multilevel data structure, as shown
// in the example below.
//
// Flatten the list so that all the nodes appear in a single-level, doubly linked list.
// You are given the head of the first level of the list.

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// Solution1: Using recursion
// We traverse the list until we find a node which has a Child reference. When we find such
// a node, we attach the child list as the Next of the current node and make the end node's
// Next of the child list point to the next node of the current node. We do this until we
// reach the end of the list.
//
// Time Complexity: O(n)	Space Complexity: O(no. of child references - 1) Implicit call stack memory
func Flatten(root *Node) *Node {
	curr := root
	for curr != nil {
		// Check if the current node has any Child reference
		if curr.Child == nil {
			// We have reached the end of the list
			curr = curr.Next
		} else {
			next := curr.Next                 // Save the current node's Next
			node := Flatten(curr.Child)       // Recall the function with the child list
			curr.Child = nil                  // Make current node's Child nil, as we will attach as the Next
			curr.Next, node.Prev = node, curr // Make current node's Next point to start of child list and start node's Prev to current node
			for curr.Next != nil {
				curr = curr.Next
			} // Traverse to end of child list
			if next == nil {
				break
			} // If Next is nil, we have reached the end, no need to point end of child list to Next
			curr.Next, next.Prev = next, curr // Make last node of child list point to Next and it's Prev to current node
			curr = curr.Next
		}
	}
	return root
}

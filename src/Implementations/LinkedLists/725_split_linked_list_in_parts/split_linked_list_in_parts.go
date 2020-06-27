package split_linked_list_in_parts

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: Given a (singly) linked list with head node root, write a function
// to split the linked list into k consecutive linked list "parts".
// The length of each part should be as equal as possible: no two parts should have a size
// differing by more than 1. This may lead to some parts being null.
// The parts should be in order of occurrence in the input list, and parts occurring earlier
// should always have a size greater than or equal parts occurring later.
// Return a List of ListNode's representing the linked list parts that are formed.

// Solution: Straight forward math
// We will split the input list directly and return a list of pointers to nodes in the original
// list as appropriate.
// Our solution proceeds similarly. For a part of size L = width + (i < remainder ? 1 : 0),
// instead of stepping L times, we will step L-1 times, and our final time will also sever the
// link between the last node from the previous part and the first node from the next part.
//
// Time Complexity: O(n + k)	Space Complexity: O(1)
func SplitListToParts(root *ListNode, k int) []*ListNode {
	nodes := 0
	curr1 := root
	for curr1 != nil {
		nodes++
		curr1 = curr1.Next
	}
	ans := make([]*ListNode, k)
	quo, rem := nodes/k, nodes%k

	curr := root
	for i := 0; curr != nil && i < k; i++ {
		ans[i] = curr
		size := quo
		if rem > 0 {
			size++
			rem--
		}
		for j := 1; j < size; j++ {
			curr = curr.Next
		}
		end := curr
		curr = curr.Next
		end.Next = nil
	}
	return ans
}

package linked_list_components

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: We are given head, the head node of a linked list containing
// unique integer values.
// We are also given the list G, a subset of the values in the linked list.
// Return the number of connected components in G, where two values are connected
// if they appear consecutively in the linked list.

// Solution: Using hashmap
// Scanning through the list, if node.val is in G and node.next.val isn't (including if node.next is null),
// then this must be the end of a connected component.
func NumComponents(head *ListNode, G []int) int {
	hash := make(map[int]bool)
	for _, val := range G {
		hash[val] = true
	}
	ans, connected := 0, false
	for curr := head; curr != nil; curr = curr.Next {
		if hash[curr.Val] {
			if !connected {
				ans++
			}
			connected = true
		} else {
			connected = false
		}
	}
	return ans
}

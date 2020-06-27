package next_greater_node_in_linked_list

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem Statement: We are given a linked list with head as the first node.
// Let's number the nodes in the list:
// node_1, node_2, node_3, ... etc.
// Each node may have a next larger value: for node_i, next_larger(node_i) is the node_j.val such that
// j > i, node_j.val > node_i.val, and j is the smallest possible choice.  If such a j
// does not exist, the next larger value is 0.
// Return an array of integers answer, where answer[i] = next_larger(node_{i+1}).
// Note that in the example inputs (not outputs) below, arrays such as [2,1,5] represent
// the serialization of a linked list with a head node value of 2, second node value of 1,
// and third node value of 5.

// Solution: Using stack
// We use a stack to keep track of the next larger element for the previous element. First,
// we convert the list to a slice and initialize an empty stack.
// For every element in the list, we check if the top of stack is less than the current element.
// If so, then for the top of the stack, the current element is the nearest larger element.
// Else, we push the element to the top of the stack. Also, if the stack is empty, we push the
// element to the stack.
//
// TimeComplexity: O(n)		Space Complexity: O(n)
func NextLargerNodes(head *ListNode) []int {
	values := listToSlice(head)
	result, stack := make([]int, len(values)), make([]int, 0)
	for i := 0; i < len(values); i++ {
		for len(stack) > 0 && values[stack[len(stack)-1]] < values[i] {
			result[stack[len(stack)-1]] = values[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}

func listToSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	curr, values := head, make([]int, 0)
	for ; curr != nil; curr = curr.Next {
		values = append(values, curr.Val)
	}
	return values
}

package linked_list_cycle_2

import . "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"

// Problem statement: Determine if a cycle exists in the given list of nodes

// Solution 1: Marking nodes
// One way of solving this problem will be marking the nodes as we visit them.
// If a node is marked and visit the same node, we can conclude that a cycle
// exists in the list. But this will require us to change the existing value.
// We mark the value as the negation of the existing value. For example we have
// three nodes
//
// 			1 --> 2 --> 3 --> 4 --> 5
//                              /|\    |
//                               -------
//
// As we visit the first node we change it to -1, then we move to the next node.
// If there is a cycle we will definitely visit the marked nodes. If not the loop
// traversing the nodes will terminate. We need to undo the negations before
// returning the function.
//
// Time Complexity: O(n)	Space Complexity: O(1)
// Note: This method will only work if all the values in the list are either positive/negative.
func HasCycle21(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	curr := head
	var ans *ListNode
	for ; curr != nil; curr = curr.Next {
		if curr.Val > 0 {
			curr.Val = -curr.Val
		} else {
			ans = curr
			break
		}
	}
	curr = head
	for ; curr != nil; curr = curr.Next {
		if curr.Val > 0 {
			break
		} else {
			curr.Val = -curr.Val
		}
	}
	return ans
}

// Solution 2: Using Floyd's algorithm
// Here you use 2 pointers, 1 moving at the speed of 1 and the other moving at the speed of 2.
// Finally the 2 pointers would meet each other if there is a cycle or the fast pointer would
// encounter null if it’s not a cycle.
//
// Time Compexity: O(n)	Space Complexity: O(1)
func HasCycle22(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	hasCycle := false
	for slow != nil && fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil
	}
	slow = head
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}

package IntersectionOfTwoLists

import . "github.com/bumblebee211996/go-ds-daa/src/Implementations/LinkedLists"

// Problem Statement: Find the node at which the intersection of two singly linked lists begins.

// Solution 1: Brute force approach
func IntersectingNode1(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	for l1 := head1; l1 != nil; l1 = l1.Next {
		for l2 := head2; l2 != nil; l2 = l2.Next {
			if l1.Val == l2.Val {
				return l1
			}
		}
	}
	return nil
}

// Solution 2: Using HashTable
// Traverse list1 and store the address/reference to each node in a hash set. Then check
// every node b in list2: if b appears in the hash set, then b is the intersection node.
func IntersectingNode2(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	hash := make(map[*ListNode]bool)
	for curr := head1; curr != nil; curr = curr.Next {
		hash[curr] = true
	}
	for curr := head2; curr != nil; curr = curr.Next {
		if _, ok := hash[curr]; ok {
			return curr
		}
	}
	return nil
}

// Solution 3: Two pointer solution
func IntersectingNode3(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	c1, c2 := head1, head2
	for c1 != c2 {
		if c1 == nil {
			c1, c2 = head1, c2.Next
		} else if c2 == nil {
			c1, c2 = c1.Next, head2
		} else {
			c1, c2 = c1.Next, c2.Next
		}
	}
	return c1
}

// Solution 4: Length solution
// Store the size of List1 and List2 as len1 and len2. Then reset the pointers to head1 and
// head2 and find the difference between len1 and len2, and then let the pointer of the longer
// list proceed by the difference between len1 and len2. Finally, traverse through the lists
// again, the intersection node can be easily found.
func IntersectingNode4(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	l1, l2 := getLength(head1), getLength(head2)
	if l1 > l2 {
		return helper(head1, head2, l1-l2)
	} else {
		return helper(head2, head1, l2-l1)
	}
}

func getLength(head *ListNode) int {
	curr, l := head, 0
	for ; curr != nil; curr = curr.Next {
		l++
	}
	return l
}

func helper(headA, headB *ListNode, d int) *ListNode {
	currA := headA
	currB := headB
	for i := 0; i < d; i++ {
		currA = currA.Next
	}
	for currA != nil && currB != nil && currA != currB {
		currA = currA.Next
		currB = currB.Next
	}
	return currA
}

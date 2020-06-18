package LinkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToListWithCycle(values []int, cycle bool) *ListNode {
	head := &ListNode{Val: values[0]}
	temp := head
	for i := 1; i < len(values); i++ {
		temp.Next = &ListNode{Val: values[i]}
		temp = temp.Next
	}
	if cycle {
		temp.Next = head.Next
	}
	return head
}

func SliceToList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	temp := head
	for i := 1; i < len(values); i++ {
		temp.Next = &ListNode{Val: values[i]}
		temp = temp.Next
	}
	return head
}

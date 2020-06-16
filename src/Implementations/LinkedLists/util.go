package LinkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToList(values []int, cycle bool) *ListNode {
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

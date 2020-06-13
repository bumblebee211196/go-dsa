package LinkedLists

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
	Len  int
}

func Copy(list *SinglyLinkedList) *SinglyLinkedList {
	newList, curr := &SinglyLinkedList{}, list.Head
	for ; curr != nil; curr = curr.Next {
		newList.AddToTail(curr.Val)
	}
	return newList
}

func SliceToList(values []int) *SinglyLinkedList {
	newList := &SinglyLinkedList{}
	for _, val := range values {
		newList.AddToTail(val)
	}
	return newList
}

func PrintList(list *SinglyLinkedList) {
	t := list.Head
	for ; t != nil; t = t.Next {
		fmt.Printf("%v --> ", t.Val)
	}
	fmt.Println()
}

func ListToSlice(list *SinglyLinkedList) []int {
	t := list.Head
	res := make([]int, 0)
	for ; t != nil; t = t.Next {
		res = append(res, t.Val)
	}
	return res
}

// Adds the node to the head
func (list *SinglyLinkedList) AddToHead(val int) {
	list.Head = &Node{Val: val, Next: list.Head}
	list.Len++
}

// Adds the node to the tail
func (list *SinglyLinkedList) AddToTail(val int) {
	if list.Head == nil {
		list.Head = &Node{Val: val}
		list.Len++
		return
	}
	curr := list.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &Node{Val: val, Next: nil}
	list.Len++
}

// Adds the node to the given index, if the index is valid
func (list *SinglyLinkedList) AddToIndex(index, val int) {
	// This function is based on 0-indexing
	if index == 0 {
		list.AddToHead(val)
	} else if index == list.Len {
		list.AddToTail(val)
	} else if index > 0 && index < list.Len {
		curr := list.Head
		for i := 0; i < index-1; i++ {
			curr = curr.Next
		}
		curr.Next = &Node{Val: val, Next: curr.Next}
		list.Len++
	}
}

// Removes the head node if there exists one
func (list *SinglyLinkedList) RemoveHead() {
	if list.Len > 0 {
		list.Head = list.Head.Next
		list.Len--
	}
}

// Removes the head node if there exists one
func (list *SinglyLinkedList) RemoveTail() {
	if list.Len > 0 {
		if list.Head.Next == nil {
			list.Head = nil
		} else {
			prev, curr := list.Head, list.Head.Next
			for ; curr.Next != nil; curr = curr.Next {
				prev, curr = prev.Next, curr.Next
			}
			prev.Next = nil
		}
		list.Len--
	}
}

// Removes the node at the given index if there exists one
func (list *SinglyLinkedList) RemoveFromIndex(index int) {
	if index == 0 {
		list.RemoveHead()
	} else if index == list.Len-1 {
		list.RemoveTail()
	} else if index > 0 && index < list.Len-1 {
		curr := list.Head
		for i := 0; i < index-1; i++ {
			curr = curr.Next
		}
		curr.Next = curr.Next.Next
		list.Len--
	}
}

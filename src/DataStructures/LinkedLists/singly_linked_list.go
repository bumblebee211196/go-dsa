package LinkedLists

// Adds the node to the head
func (list *SinglyLinkedList) AddToHead(val int) {
	list.Head = &Node{Val: val, Next: list.Head}
	list.Len++
}

// Adds the node to the given index, if the index is valid
func (list *SinglyLinkedList) AddToIndex(index, val int) {
	// This function is based on 0-indexing
	if index == 0 {
		list.AddToHead(val)
	} else if index > 0 && index <= list.Len {
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

// Removes the node at the given index if there exists one
func (list *SinglyLinkedList) RemoveFromIndex(index int) {
	if index == 0 {
		list.RemoveHead()
	} else if index > 0 && index <= list.Len-1 {
		curr := list.Head
		for i := 0; i < index-1; i++ {
			curr = curr.Next
		}
		curr.Next = curr.Next.Next
		list.Len--
	}
}

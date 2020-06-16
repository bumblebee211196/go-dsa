package LinkedLists

// Adds the node to the head
func (list *DoublyLinkedList) AddToHead(val int) {
	list.Head = &Node{Val: val, Next: list.Head, Prev: nil}
	list.Len++
}

// Adds the node to the given index, if the index is valid
func (list *DoublyLinkedList) AddToIndex(index, val int) {
	if index == 0 {
		list.AddToHead(val)
	} else if index > 0 && index <= list.Len {
		curr := list.Head
		for i := 0; i < index-1; i++ {
			curr = curr.Next
		}
		next := curr.Next
		curr.Next = &Node{Val: val, Prev: curr, Next: next}
		if next != nil {
			next.Prev = curr.Next
		}
		list.Len++
	}
}

// Removes the head node if there exists one
func (list *DoublyLinkedList) RemoveHead() {
	if list.Len > 0 {
		list.Head = list.Head.Next
		if list.Head != nil {
			list.Head.Prev = nil
		}
		list.Len--
	}
}

// Removes the node at the given index if there exists one
func (list *DoublyLinkedList) RemoveFromIndex(index int) {
	if index == 0 {
		list.RemoveHead()
	} else if index > 0 && index <= list.Len-1 {
		curr := list.Head
		for i := 0; i < index-1; i++ {
			curr = curr.Next
		}
		next := curr.Next.Next
		if next != nil {
			curr.Next, next.Prev = next, curr
		} else {
			curr.Next = next
		}
		list.Len--
	}
}

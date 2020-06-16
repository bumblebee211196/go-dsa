package LinkedLists

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
	Len  int
}

type DoublyLinkedList struct {
	Head *Node
	Len  int
}

func Copy(expr interface{}) interface{} {
	switch expr.(type) {
	case SinglyLinkedList:
		list := expr.(SinglyLinkedList)
		newList, curr := &SinglyLinkedList{}, list.Head
		for ; curr != nil; curr = curr.Next {
			newList.AddToHead(curr.Val)
		}
		return newList
	case DoublyLinkedList:
		list := expr.(DoublyLinkedList)
		newList, curr := &DoublyLinkedList{}, list.Head
		for ; curr != nil; curr = curr.Next {
			newList.AddToHead(curr.Val)
		}
		return newList
	default:
		return nil
	}
}

func SliceToList(values []int, expr string) interface{} {
	switch expr {
	case "SinglyLinkedList":
		newList := &SinglyLinkedList{}
		for i := len(values) - 1; i >= 0; i-- {
			newList.AddToHead(values[i])
		}
		return newList
	case "DoublyLinkedList":
		newList := &DoublyLinkedList{}
		for i := len(values) - 1; i >= 0; i-- {
			newList.AddToHead(values[i])
		}
		return newList
	default:
		return nil
	}
}

func ListToSlice(expr interface{}) []int {
	switch expr.(type) {
	case SinglyLinkedList:
		list := expr.(SinglyLinkedList)
		curr, res := list.Head, make([]int, 0)
		for ; curr != nil; curr = curr.Next {
			res = append(res, curr.Val)
		}
		return res
	case DoublyLinkedList:
		list := expr.(DoublyLinkedList)
		curr, res := list.Head, make([]int, 0)
		for ; curr != nil; curr = curr.Next {
			res = append(res, curr.Val)
		}
		return res
	default:
		return nil
	}
}

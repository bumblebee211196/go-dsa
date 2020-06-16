package LinkedLists

import (
	"reflect"
	"strconv"
	"testing"
)

func TestDoublyLinkedList_AddToHead(t *testing.T) {
	cases := []struct {
		name   string
		list   *DoublyLinkedList
		val    int
		expect *DoublyLinkedList
	}{
		{"TestCase", SliceToList([]int{}, "DoublyLinkedList").(*DoublyLinkedList), 5, SliceToList([]int{5}, "DoublyLinkedList").(*DoublyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList), 5, SliceToList([]int{5, 1, 2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList)},
	}
	for i, c := range cases {
		t.Run(c.name+strconv.Itoa(i+1), func(t *testing.T) {
			input := Copy(c.list)
			c.list.AddToHead(c.val)
			if !reflect.DeepEqual(ListToSlice(c.list), ListToSlice(c.expect)) {
				t.Fatalf("Expected %v, got %v for inputs: %v %v",
					c.expect, c.list, input, c.val)
			}
		})
	}
}

func TestDoublyLinkedList_AddToIndex(t *testing.T) {
	cases := []struct {
		name   string
		list   *DoublyLinkedList
		index  int
		val    int
		expect *DoublyLinkedList
	}{
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList), 0, 5, SliceToList([]int{5, 1, 2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList), 1, 5, SliceToList([]int{1, 5, 2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList), 2, 5, SliceToList([]int{1, 2, 5, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList), 3, 5, SliceToList([]int{1, 2, 3, 5, 4}, "DoublyLinkedList").(*DoublyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList), 4, 5, SliceToList([]int{1, 2, 3, 4, 5}, "DoublyLinkedList").(*DoublyLinkedList)},
	}
	for i, c := range cases {
		t.Run(c.name+strconv.Itoa(i+1), func(t *testing.T) {
			input := Copy(c.list)
			c.list.AddToIndex(c.index, c.val)
			if !reflect.DeepEqual(ListToSlice(c.list), ListToSlice(c.expect)) {
				t.Fatalf("Expected %v, got %v for inputs: %v %v %v",
					c.expect, c.list, input, c.index, c.val)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveHead(t *testing.T) {
	cases := []struct {
		name   string
		list   *DoublyLinkedList
		expect *DoublyLinkedList
	}{
		{"TestCase", SliceToList([]int{}, "DoublyLinkedList").(*DoublyLinkedList), SliceToList([]int{}, "DoublyLinkedList").(*DoublyLinkedList)},
		{"TestCase", SliceToList([]int{1}, "DoublyLinkedList").(*DoublyLinkedList), SliceToList([]int{}, "DoublyLinkedList").(*DoublyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList), SliceToList([]int{2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList)},
	}
	for i, c := range cases {
		t.Run(c.name+strconv.Itoa(i+1), func(t *testing.T) {
			input := Copy(c.list)
			c.list.RemoveHead()
			if !reflect.DeepEqual(ListToSlice(c.list), ListToSlice(c.expect)) {
				t.Fatalf("Expected %v, got %v for inputs: %v",
					c.expect, c.list, input)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveFromIndex(t *testing.T) {
	cases := []struct {
		name   string
		list   *DoublyLinkedList
		index  int
		expect *DoublyLinkedList
	}{
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}, "DoublyLinkedList").(*DoublyLinkedList), 0, SliceToList([]int{2, 3, 4, 5}, "DoublyLinkedList").(*DoublyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}, "DoublyLinkedList").(*DoublyLinkedList), 1, SliceToList([]int{1, 3, 4, 5}, "DoublyLinkedList").(*DoublyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}, "DoublyLinkedList").(*DoublyLinkedList), 2, SliceToList([]int{1, 2, 4, 5}, "DoublyLinkedList").(*DoublyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}, "DoublyLinkedList").(*DoublyLinkedList), 3, SliceToList([]int{1, 2, 3, 5}, "DoublyLinkedList").(*DoublyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}, "DoublyLinkedList").(*DoublyLinkedList), 4, SliceToList([]int{1, 2, 3, 4}, "DoublyLinkedList").(*DoublyLinkedList)},
	}
	for i, c := range cases {
		t.Run(c.name+strconv.Itoa(i+1), func(t *testing.T) {
			input := Copy(c.list)
			c.list.RemoveFromIndex(c.index)
			if !reflect.DeepEqual(ListToSlice(c.list), ListToSlice(c.expect)) {
				t.Fatalf("Expected %v, got %v for inputs: %v %v",
					c.expect, c.list, input, c.index)
			}
		})
	}
}

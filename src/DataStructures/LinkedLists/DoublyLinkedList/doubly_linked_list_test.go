package DoublyLinkedList

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
		{"TestCase", SliceToList([]int{}), 10, SliceToList([]int{10})},
		{"TestCase", SliceToList([]int{6, 7, 8, 9}), 5, SliceToList([]int{5, 6, 7, 8, 9})},
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
		{"TestCase", SliceToList([]int{6, 7, 8, 9}), 0, 5, SliceToList([]int{5, 6, 7, 8, 9})},
		{"TestCase", SliceToList([]int{6, 7, 8, 9}), 1, 5, SliceToList([]int{6, 5, 7, 8, 9})},
		{"TestCase", SliceToList([]int{6, 7, 8, 9}), 2, 5, SliceToList([]int{6, 7, 5, 8, 9})},
		{"TestCase", SliceToList([]int{6, 7, 8, 9}), 3, 5, SliceToList([]int{6, 7, 8, 5, 9})},
		{"TestCase", SliceToList([]int{6, 7, 8, 9}), 4, 5, SliceToList([]int{6, 7, 8, 9, 5})},
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
		{"TestCase", SliceToList([]int{}), SliceToList([]int{})},
		{"TestCase", SliceToList([]int{6}), SliceToList([]int{})},
		{"TestCase", SliceToList([]int{6, 7, 8, 9}), SliceToList([]int{2, 3, 4})},
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
		{"TestCase", SliceToList([]int{6, 7, 8, 9, 10}), 0, SliceToList([]int{7, 8, 9, 10})},
		{"TestCase", SliceToList([]int{6, 7, 8, 9, 10}), 1, SliceToList([]int{6, 8, 9, 10})},
		{"TestCase", SliceToList([]int{6, 7, 8, 9, 10}), 2, SliceToList([]int{6, 7, 9, 10})},
		{"TestCase", SliceToList([]int{6, 7, 8, 9, 10}), 3, SliceToList([]int{6, 7, 8, 10})},
		{"TestCase", SliceToList([]int{6, 7, 8, 9, 10}), 4, SliceToList([]int{6, 7, 8, 9})},
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

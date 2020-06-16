package LinkedLists

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSinglyLinkedList_AddToHead(t *testing.T) {
	cases := []struct {
		name   string
		list   *SinglyLinkedList
		val    int
		expect *SinglyLinkedList
	}{
		{"TestCase", SliceToList([]int{}, "SinglyLinkedList").(*SinglyLinkedList), 5, SliceToList([]int{5}, "SinglyLinkedList").(*SinglyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList), 5, SliceToList([]int{5, 1, 2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList)},
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

func TestSinglyLinkedList_AddToIndex(t *testing.T) {
	cases := []struct {
		name   string
		list   *SinglyLinkedList
		index  int
		val    int
		expect *SinglyLinkedList
	}{
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList), 0, 5, SliceToList([]int{5, 1, 2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList), 1, 5, SliceToList([]int{1, 5, 2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList), 2, 5, SliceToList([]int{1, 2, 5, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList), 3, 5, SliceToList([]int{1, 2, 3, 5, 4}, "SinglyLinkedList").(*SinglyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList), 4, 5, SliceToList([]int{1, 2, 3, 4, 5}, "SinglyLinkedList").(*SinglyLinkedList)},
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

func TestSinglyLinkedList_RemoveHead(t *testing.T) {
	cases := []struct {
		name   string
		list   *SinglyLinkedList
		expect *SinglyLinkedList
	}{
		{"TestCase", SliceToList([]int{}, "SinglyLinkedList").(*SinglyLinkedList), SliceToList([]int{}, "SinglyLinkedList").(*SinglyLinkedList)},
		{"TestCase", SliceToList([]int{1}, "SinglyLinkedList").(*SinglyLinkedList), SliceToList([]int{}, "SinglyLinkedList").(*SinglyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList), SliceToList([]int{2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList)},
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

func TestSinglyLinkedList_RemoveFromIndex(t *testing.T) {
	cases := []struct {
		name   string
		list   *SinglyLinkedList
		index  int
		expect *SinglyLinkedList
	}{
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}, "SinglyLinkedList").(*SinglyLinkedList), 0, SliceToList([]int{2, 3, 4, 5}, "SinglyLinkedList").(*SinglyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}, "SinglyLinkedList").(*SinglyLinkedList), 1, SliceToList([]int{1, 3, 4, 5}, "SinglyLinkedList").(*SinglyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}, "SinglyLinkedList").(*SinglyLinkedList), 2, SliceToList([]int{1, 2, 4, 5}, "SinglyLinkedList").(*SinglyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}, "SinglyLinkedList").(*SinglyLinkedList), 3, SliceToList([]int{1, 2, 3, 5}, "SinglyLinkedList").(*SinglyLinkedList)},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}, "SinglyLinkedList").(*SinglyLinkedList), 4, SliceToList([]int{1, 2, 3, 4}, "SinglyLinkedList").(*SinglyLinkedList)},
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

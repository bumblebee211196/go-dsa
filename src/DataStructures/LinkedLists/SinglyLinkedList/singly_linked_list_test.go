package SinglyLinkedList

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
		{"TestCase", SliceToList([]int{}), 5, SliceToList([]int{5})},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}), 5, SliceToList([]int{5, 1, 2, 3, 4})},
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
		{"TestCase", SliceToList([]int{1, 2, 3, 4}), 0, 5, SliceToList([]int{5, 1, 2, 3, 4})},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}), 1, 5, SliceToList([]int{1, 5, 2, 3, 4})},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}), 2, 5, SliceToList([]int{1, 2, 5, 3, 4})},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}), 3, 5, SliceToList([]int{1, 2, 3, 5, 4})},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}), 4, 5, SliceToList([]int{1, 2, 3, 4, 5})},
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
		{"TestCase", SliceToList([]int{}), SliceToList([]int{})},
		{"TestCase", SliceToList([]int{1}), SliceToList([]int{})},
		{"TestCase", SliceToList([]int{1, 2, 3, 4}), SliceToList([]int{2, 3, 4})},
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
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}), 0, SliceToList([]int{2, 3, 4, 5})},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}), 1, SliceToList([]int{1, 3, 4, 5})},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}), 2, SliceToList([]int{1, 2, 4, 5})},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}), 3, SliceToList([]int{1, 2, 3, 5})},
		{"TestCase", SliceToList([]int{1, 2, 3, 4, 5}), 4, SliceToList([]int{1, 2, 3, 4})},
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

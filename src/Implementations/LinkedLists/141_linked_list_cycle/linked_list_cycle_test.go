package linked_list_cycle

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestHasCycle1(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect bool
	}{
		{SliceToListWithCycle([]int{1}, false), false},
		{SliceToListWithCycle([]int{1, 2, 3, 4, 5}, true), true},
		{SliceToListWithCycle([]int{1, 2, 3, 4, 5}, false), false},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := HasCycle1(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					c.expect, got, ListToSlice(c.input))
			}
		})
	}
}

func TestHasCycle2(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect bool
	}{
		{SliceToListWithCycle([]int{1}, false), false},
		{SliceToListWithCycle([]int{1, 2, 3, 4, 5}, true), true},
		{SliceToListWithCycle([]int{1, 2, 3, 4, 5}, false), false},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := HasCycle2(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					c.expect, got, ListToSlice(c.input))
			}
		})
	}
}

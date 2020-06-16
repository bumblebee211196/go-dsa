package LinkedLists

import (
	"reflect"
	"strconv"
	"testing"
)

func TestHasCycle21(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{1}, false), nil},
		{SliceToList([]int{1, 2, 3, 4, 5}, true), &ListNode{Val: 2}},
		{SliceToList([]int{1, 2, 3, 4, 5}, false), nil},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := HasCycle21(c.input)
			if got != nil && c.expect != nil && !reflect.DeepEqual(got.Val, c.expect.Val) {
				t.Fatalf("Expected: %v, got %v for input %v",
					c.expect, got, c.input)
			}
		})
	}
}

func TestHasCycle22(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{1}, false), nil},
		{SliceToList([]int{1, 2, 3, 4, 5}, true), &ListNode{Val: 2}},
		{SliceToList([]int{1, 2, 3, 4, 5}, false), nil},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := HasCycle22(c.input)
			if got != nil && c.expect != nil && !reflect.DeepEqual(got.Val, c.expect.Val) {
				t.Fatalf("Expected: %v, got %v for input %v",
					c.expect, got, c.input)
			}
		})
	}
}
